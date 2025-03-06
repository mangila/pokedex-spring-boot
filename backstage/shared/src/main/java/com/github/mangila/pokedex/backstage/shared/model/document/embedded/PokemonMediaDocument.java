package com.github.mangila.pokedex.backstage.shared.model.document.embedded;

import com.github.mangila.pokedex.backstage.model.grpc.mongodb.PokemonMediaPrototype;

public record PokemonMediaDocument() {

    public PokemonMediaPrototype toProto() {
        return PokemonMediaPrototype.newBuilder().build();
    }

    public static PokemonMediaDocument fromProto(PokemonMediaPrototype pokemonMediaPrototype) {
        return new PokemonMediaDocument();
    }
}
