package com.github.mangila.pokedex.backstage.shared.model.domain;

public enum RedisConsumerGroup {

    POKEDEX_BACKSTAGE_GROUP("pokedex-backstage-group");

    private final String groupName;

    RedisConsumerGroup(String groupName) {
        this.groupName = groupName;
    }

    public String getGroupName() {
        return groupName;
    }
}

