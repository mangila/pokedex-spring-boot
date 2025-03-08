package com.github.mangila.pokedex.backstage.bouncer.pokeapi.http.response.pokemon;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.github.mangila.pokedex.backstage.model.grpc.pokeapi.CriesPrototype;
import org.apache.logging.log4j.util.Strings;

import java.util.Objects;

public record Cries(
        @JsonProperty("latest") String latest,
        @JsonProperty("legacy") String legacy
) {
    public CriesPrototype toProto() {
        return CriesPrototype.newBuilder()
                .setLatest(Objects.nonNull(latest) ? latest : Strings.EMPTY)
                .setLegacy(Objects.nonNull(legacy) ? legacy : Strings.EMPTY)
                .build();
    }
}
