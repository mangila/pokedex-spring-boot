package com.github.mangila.pokedex.backstage.generation.task;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.github.mangila.pokedex.backstage.integration.bouncer.redis.RedisBouncerClient;
import com.github.mangila.pokedex.backstage.integration.pokeapi.PokeApiTemplate;
import com.github.mangila.pokedex.backstage.integration.pokeapi.response.generation.GenerationResponse;
import com.github.mangila.pokedex.backstage.integration.pokeapi.response.generation.Species;
import com.github.mangila.pokedex.backstage.model.Generation;
import com.github.mangila.pokedex.backstage.model.PokemonName;
import com.github.mangila.pokedex.backstage.model.RedisStreamKey;
import com.github.mangila.pokedex.backstage.model.Task;
import com.github.mangila.pokedex.backstage.model.grpc.redis.StreamRecord;
import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.util.EnumSet;
import java.util.List;

@Service
public class GenerationTask implements Task {

    private static final Logger log = LoggerFactory.getLogger(GenerationTask.class);

    private final PokeApiTemplate pokeApiTemplate;
    private final RedisBouncerClient redisBouncerClient;
    private final ObjectMapper objectMapper;

    public GenerationTask(PokeApiTemplate pokeApiTemplate,
                          RedisBouncerClient redisBouncerClient,
                          ObjectMapper objectMapper) {
        this.pokeApiTemplate = pokeApiTemplate;
        this.redisBouncerClient = redisBouncerClient;
        this.objectMapper = objectMapper;
    }

    /**
     * 0. Create a grpc client-stream to RedisBouncer and add Redis Stream messages
     * 1. Iterate all Generation enums
     * 2. Map to Generation name
     * 3. Check if generation is in Redis Cache else send Api request
     * 4. Flatten out response list with Pokemon Species
     * 5. Create a PokemonName object and Validate
     * 6. Map to json string
     * 7. Put all Pokemons on GENERATION_STREAM
     *
     * @param args - program arguments from the Main method
     */
    @Override
    public void run(String[] args) {
        var observer = redisBouncerClient.streamOps()
                .addWithClientStream(new StreamObserver<>() {
                    @Override
                    public void onNext(Empty empty) {

                    }

                    @Override
                    public void onError(Throwable throwable) {

                    }

                    @Override
                    public void onCompleted() {

                    }
                });
        EnumSet.allOf(Generation.class)
                .stream()
                .map(Generation::getName)
                .peek(generation -> log.info("Generation push to Queue: {}", generation))
                .map(generationName -> {
                    var cacheValue = redisBouncerClient.valueOps().get(generationName);
                    if (cacheValue.isEmpty()) {
                        var response = pokeApiTemplate.fetchGeneration(generationName);
                        redisBouncerClient.valueOps()
                                .set(generationName, response.toJson(objectMapper));
                        return response;
                    }
                    try {
                        return objectMapper.readValue(cacheValue.get(), GenerationResponse.class);
                    } catch (JsonProcessingException e) {
                        throw new RuntimeException(e);
                    }
                })
                .map(GenerationResponse::pokemonSpecies)
                .flatMap(List::stream)
                .map(Species::name)
                .map(PokemonName::new)
                .forEach(pokemonName -> {
                    var record = StreamRecord.newBuilder()
                            .setStreamKey(RedisStreamKey.POKEMON_NAME_EVENT.getKey())
                            .putData("name", pokemonName.name())
                            .build();
                    observer.onNext(record);
                });
        observer.onCompleted();
    }
}
