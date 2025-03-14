package com.github.mangila.pokedex.backstage.media.task;

import org.testcontainers.containers.GenericContainer;
import org.testcontainers.containers.MongoDBContainer;
import org.testcontainers.containers.Network;
import org.testcontainers.containers.wait.strategy.LogMessageWaitStrategy;
import org.testcontainers.utility.DockerImageName;

import java.time.Duration;
import java.util.List;

final class TestContainerUtil {

    private TestContainerUtil() {
        throw new IllegalStateException("Utility class");
    }

    private static final DockerImageName IMAGE_CONVERTER_NAME = DockerImageName.parse("mangila/pokedex-image-converter");
    private static final DockerImageName REDIS_CONTAINER_NAME = DockerImageName.parse("redis:7.4.2-alpine");
    private static final DockerImageName MONGODB_IMAGE_NAME = DockerImageName.parse("mongo");
    private static final DockerImageName REDIS_BOUNCER_CONTAINER_NAME = DockerImageName.parse("mangila/pokedex-redis-bouncer");
    private static final DockerImageName MONGO_DB_BOUNCER_CONTAINER_NAME = DockerImageName.parse("mangila/pokedex-mongodb-bouncer");

    @SuppressWarnings({"rawtypes"})
    static GenericContainer<?> buildRedis() {
        return new GenericContainer(REDIS_CONTAINER_NAME)
                .withNetworkAliases("redis")
                .withNetwork(Network.SHARED);
    }

    @SuppressWarnings({"unchecked", "rawtypes"})
    static GenericContainer<?> buildRedisBouncer(String port) {
        var redisBouncer = new GenericContainer(REDIS_BOUNCER_CONTAINER_NAME)
                .withNetworkAliases("redis-bouncer")
                .withNetwork(Network.SHARED)
                .withEnv("spring.grpc.server.port", port)
                .withEnv("spring.data.redis.host", "redis")
                .withEnv("spring.data.redis.port", "6379")
                .waitingFor(new LogMessageWaitStrategy()
                        .withRegEx(".*gRPC Server started.*")
                        .withTimes(1)
                        .withStartupTimeout(Duration.ofSeconds(1)));
        redisBouncer.setPortBindings(List.of(
                port.concat(":").concat(port) //host:container port
        ));
        return redisBouncer;
    }

    static MongoDBContainer buildMongoDb() {
        return new MongoDBContainer(MONGODB_IMAGE_NAME)
                .withNetworkAliases("mongo")
                .withNetwork(Network.SHARED);
    }

    @SuppressWarnings({"rawtypes", "unchecked"})
    static GenericContainer<?> buildMongoDbBouncer(String port) {
        var mongoDbBouncer = new GenericContainer(MONGO_DB_BOUNCER_CONTAINER_NAME)
                .withNetwork(Network.SHARED)
                .withEnv("spring.grpc.server.port", port)
                .withEnv("spring.data.mongodb.uri", "mongodb://mongo:27017")
                .waitingFor(new LogMessageWaitStrategy()
                        .withRegEx(".*gRPC Server started.*")
                        .withTimes(1)
                        .withStartupTimeout(Duration.ofSeconds(1)));
        mongoDbBouncer.setPortBindings(List.of(
                port.concat(":").concat(port) //host:container port
        ));
        return mongoDbBouncer;
    }

    @SuppressWarnings({"unchecked", "rawtypes"})
    public static GenericContainer<?> buildImageConverter(String port) {
        var imageConverter = new GenericContainer(IMAGE_CONVERTER_NAME)
                .withNetwork(Network.SHARED)
                .withEnv("PORT", port)
                .waitingFor(new LogMessageWaitStrategy()
                        .withRegEx(".*server listening at.*")
                        .withTimes(1)
                        .withStartupTimeout(Duration.ofSeconds(1)));
        imageConverter.setPortBindings(List.of(
                port.concat(":").concat(port) //host:container port
        ));
        return imageConverter;
    }
}
