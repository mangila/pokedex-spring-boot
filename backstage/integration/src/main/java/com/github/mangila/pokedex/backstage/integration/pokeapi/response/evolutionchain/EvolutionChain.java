package com.github.mangila.pokedex.backstage.integration.pokeapi.response.evolutionchain;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public record EvolutionChain(
        @JsonProperty("evolves_to")
        List<EvolutionChain> nextChain,
        @JsonProperty("species")
        Species species
) {
}
