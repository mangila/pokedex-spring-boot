package com.github.mangila.pokedex.backstage.integration.bouncer.redis;

import com.github.mangila.pokedex.backstage.model.grpc.redis.EntryRequest;
import com.google.protobuf.Empty;
import com.google.protobuf.Message;

import java.util.Optional;

public interface ValueOperation {

    Empty set(EntryRequest request);

    <T extends Message> Optional<T> get(EntryRequest request, Class<T> clazz);
}
