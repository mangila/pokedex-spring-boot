package com.github.mangila.pokedex.backstage.bouncer.mongodb.service;

import com.github.mangila.pokedex.backstage.model.grpc.mongodb.MongoDbOperationGrpc;
import com.github.mangila.pokedex.backstage.model.grpc.mongodb.PokemonSpeciesPrototype;
import com.github.mangila.pokedex.backstage.shared.model.document.PokemonSpeciesDocument;
import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.grpc.server.service.GrpcService;

@GrpcService
public class MongoService extends MongoDbOperationGrpc.MongoDbOperationImplBase {

    private final MongoTemplate mongoTemplate;

    public MongoService(MongoTemplate mongoTemplate) {
        this.mongoTemplate = mongoTemplate;
    }

    @Override
    public void saveOne(PokemonSpeciesPrototype request, StreamObserver<Empty> responseObserver) {
        var document = PokemonSpeciesDocument.fromProto(request);
        mongoTemplate.save(document);
        responseObserver.onNext(Empty.getDefaultInstance());
        responseObserver.onCompleted();
    }
}
