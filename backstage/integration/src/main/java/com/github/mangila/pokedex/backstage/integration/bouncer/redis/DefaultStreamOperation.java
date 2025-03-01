package com.github.mangila.pokedex.backstage.integration.bouncer.redis;

import com.github.mangila.pokedex.backstage.model.grpc.redis.StreamOperationGrpc;
import com.github.mangila.pokedex.backstage.model.grpc.redis.StreamRecord;
import com.google.protobuf.Empty;
import io.grpc.stub.StreamObserver;

class DefaultStreamOperation implements StreamOperation {

    private final StreamOperationGrpc.StreamOperationBlockingStub streamOperationBlockingStub;
    private final StreamOperationGrpc.StreamOperationStub streamOperationStub;

    public DefaultStreamOperation(StreamOperationGrpc.StreamOperationBlockingStub streamOperationBlockingStub,
                                  StreamOperationGrpc.StreamOperationStub streamOperationStub) {
        this.streamOperationBlockingStub = streamOperationBlockingStub;
        this.streamOperationStub = streamOperationStub;
    }

    @Override
    public StreamRecord readOne(StreamRecord request) {
        return streamOperationBlockingStub.readOne(request);
    }

    @Override
    public StreamObserver<StreamRecord> addWithClientStream(StreamObserver<Empty> responseObserver) {
        return streamOperationStub.addWithClientStream(responseObserver);
    }
}
