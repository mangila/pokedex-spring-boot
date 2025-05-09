package com.github.mangila.pokedex.shared.pokeapi.response.evolutionchain;

import com.fasterxml.jackson.annotation.JsonProperty;

import java.util.List;

public record Chain(
        @JsonProperty("evolves_to")
        List<EvolutionChain> firstChain,
        @JsonProperty("species")
        Species species
) {
}
