package com.github.mangila.pokedex.scheduler.pokeapi;

import lombok.extern.slf4j.Slf4j;
import org.springframework.data.redis.core.RedisTemplate;
import org.springframework.http.HttpHeaders;
import org.springframework.http.MediaType;
import org.springframework.http.client.ReactorClientHttpRequestFactory;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

import java.net.URI;
import java.util.Objects;

@Slf4j
@Service
public class PokeApiTemplate {

    private final RestClient http;
    private final RedisTemplate<String, Object> redisTemplate;

    public PokeApiTemplate(RedisTemplate<String, Object> redisTemplate) {
        this.http = RestClient.builder()
                .requestFactory(new ReactorClientHttpRequestFactory())
                .defaultHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE)
                .defaultHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE)
                .defaultHeader(HttpHeaders.USER_AGENT, "PokeApiTemplate")
                .defaultHeader(HttpHeaders.ACCEPT_ENCODING, "gzip")
                .build();
        this.redisTemplate = redisTemplate;
    }

    public <T> T fetchByUrl(URI uri, Class<T> clazz) {
        ensureUriFromPokeApi(uri);
        var cacheValue = redisTemplate.opsForValue().get(uri.toString());
        if (Objects.nonNull(cacheValue)) {
            log.debug("Cache hit - {}", cacheValue);
            return clazz.cast(cacheValue);
        }
        log.debug("Cache miss");
        var response = http.get()
                .uri(uri)
                .retrieve()
                .body(clazz);
        if (Objects.nonNull(response)) {
            redisTemplate.opsForValue().set(uri.toString(), response);
        }
        return response;
    }

    private void ensureUriFromPokeApi(URI uri) {
        if (!uri.getHost().startsWith("pokeapi.co")) {
            throw new IllegalArgumentException("PokeApiTemplate only supports hosts of 'pokeapi.com'");
        }
    }
}
