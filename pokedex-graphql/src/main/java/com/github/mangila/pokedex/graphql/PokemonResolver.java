package com.github.mangila.pokedex.graphql;

import lombok.AllArgsConstructor;
import org.springframework.data.mongodb.core.MongoTemplate;
import org.springframework.data.mongodb.core.query.Criteria;
import org.springframework.data.mongodb.core.query.Query;
import org.springframework.graphql.data.method.annotation.Argument;
import org.springframework.graphql.data.method.annotation.QueryMapping;
import org.springframework.stereotype.Controller;

@Controller
@AllArgsConstructor
public class PokemonResolver {

    private final MongoTemplate mongoTemplate;

    @QueryMapping
    public Object findById(@Argument Integer id) {
        return mongoTemplate.findById(id, Object.class, "pokemon");
    }

    @QueryMapping
    public Object findByName(@Argument String name) {
        return mongoTemplate.findOne(Query.query(Criteria.where("name").is(name)), Object.class, "pokemon");
    }
}
