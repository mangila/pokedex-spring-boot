package com.github.mangila.pokedex.backstage.bouncer.pokeapi.http.response.pokemon.sprites;

import com.fasterxml.jackson.annotation.JsonProperty;

public record GenerationVII(
        @JsonProperty("icons") Icons icons,
        @JsonProperty("ultra_sun_ultra_moon") UltraSunUltraMoon ultraSunUltraMoon
) {}
