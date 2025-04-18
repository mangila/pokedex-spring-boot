package com.github.mangila.pokedex.scheduler.service;

import com.github.mangila.pokedex.scheduler.domain.PokemonEntry;
import com.github.mangila.pokedex.scheduler.pokeapi.PokeApiTemplate;
import com.github.mangila.pokedex.scheduler.pokeapi.response.evolutionchain.EvolutionChainResponse;
import com.github.mangila.pokedex.scheduler.pokeapi.response.pokemon.PokemonResponse;
import com.github.mangila.pokedex.scheduler.pokeapi.response.species.SpeciesResponse;
import com.github.mangila.pokedex.scheduler.repository.PokemonSpeciesRepository;
import com.github.mangila.pokedex.scheduler.repository.document.PokemonSpeciesDocument;
import org.springframework.data.util.Pair;
import org.springframework.stereotype.Service;

import java.net.URI;

@Service
@lombok.AllArgsConstructor
@lombok.extern.slf4j.Slf4j
public class PokemonTask {

    private final PokeApiTemplate pokeApiTemplate;
    private final PokemonSpeciesRepository pokemonSpeciesRepository;
    private final PokemonMediaHandler pokemonMediaHandler;

    public void run(PokemonEntry pokemonEntry) {
        var speciesResponse = pokeApiTemplate.fetchByUrl(pokemonEntry.uri(), SpeciesResponse.class);
        var evolutionChainResponse = pokeApiTemplate.fetchByUrl(URI.create(speciesResponse.evolutionChain().url()), EvolutionChainResponse.class);
        var varieties = speciesResponse.varieties()
                .stream()
                .peek(variety -> log.debug("Processing varieties: {} - {}", variety.pokemon().name(), pokemonEntry.name()))
                .map(variety -> variety.pokemon().url())
                .map(URI::create)
                .map(uri -> pokeApiTemplate.fetchByUrl(uri, PokemonResponse.class))
                .toList();
        var document = PokemonSpeciesDocument.of(speciesResponse, evolutionChainResponse, varieties);
        pokemonSpeciesRepository.save(document);
        varieties.forEach(variety -> pokemonMediaHandler.queueMedia(Pair.of(speciesResponse.id(), variety.id()),
                variety.name(),
                variety.sprites(),
                variety.cries()));
    }

}
