package com.github.mangila.pokedex.backstage.bouncer.pokeapi.http.response.pokemon.sprites;

import com.fasterxml.jackson.annotation.JsonProperty;

public record FireredLeafgreen(
        @JsonProperty("back_default") String backDefault,
        @JsonProperty("back_shiny") String backShiny,
        @JsonProperty("front_default") String frontDefault,
        @JsonProperty("front_shiny") String frontShiny
) {}
