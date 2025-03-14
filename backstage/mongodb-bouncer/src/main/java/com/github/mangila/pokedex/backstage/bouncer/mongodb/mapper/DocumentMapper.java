package com.github.mangila.pokedex.backstage.bouncer.mongodb.mapper;

import com.github.mangila.pokedex.backstage.bouncer.mongodb.document.PokemonSpeciesDocument;
import com.github.mangila.pokedex.backstage.bouncer.mongodb.document.embedded.*;
import com.github.mangila.pokedex.backstage.bouncer.mongodb.props.MongoDbBouncerFileApiProperties;
import com.github.mangila.pokedex.backstage.model.grpc.model.*;
import org.springframework.stereotype.Component;
import org.springframework.web.util.UriComponentsBuilder;

import java.util.Collections;
import java.util.List;

@Component
public class DocumentMapper {

    private final MongoDbBouncerFileApiProperties fileApiProperties;

    public DocumentMapper(MongoDbBouncerFileApiProperties fileApiProperties) {
        this.fileApiProperties = fileApiProperties;
    }

    public PokemonSpeciesDocument toDocument(PokemonSpecies request) {
        return new PokemonSpeciesDocument(
                request.getSpeciesId(),
                request.getName(),
                request.getGeneration(),
                toNames(request.getNamesList()),
                toDescriptions(request.getDescriptionsList()),
                toGenera(request.getGeneraList()),
                toEvolutions(request.getEvolutionsList()),
                toVarieties(request.getVarietiesList()),
                toSpecial(request.getSpecial())
        );
    }

    private List<PokemonNameDocument> toNames(List<PokemonName> namesList) {
        return namesList.stream()
                .map(pokemonName -> new PokemonNameDocument(pokemonName.getName(), pokemonName.getLanguage()))
                .toList();
    }

    private List<PokemonDescriptionDocument> toDescriptions(List<PokemonDescription> descriptionsList) {
        return descriptionsList.stream()
                .map(description -> new PokemonDescriptionDocument(description.getDescription(), description.getLanguage()))
                .toList();
    }

    private List<PokemonGeneraDocument> toGenera(List<PokemonGenera> generaList) {
        return generaList.stream()
                .map(genera -> new PokemonGeneraDocument(genera.getGenera(), genera.getLanguage()))
                .toList();
    }

    private List<PokemonEvolutionDocument> toEvolutions(List<PokemonEvolution> evolutionsList) {
        return evolutionsList.stream()
                .map(evolution -> new PokemonEvolutionDocument(evolution.getOrder(), evolution.getName()))
                .toList();
    }

    private List<PokemonDocument> toVarieties(List<Pokemon> varietiesList) {
        return varietiesList.stream()
                .map(variety -> new PokemonDocument(
                        variety.getPokemonId(),
                        variety.getName(),
                        variety.getIsDefault(),
                        variety.getHeight(),
                        variety.getWeight(),
                        variety.getTypesList().stream().map(PokemonTypeDocument::new).toList(),
                        toStats(variety.getStatsList()),
                        Collections.emptyList()
                ))
                .toList();
    }

    private List<PokemonStatDocument> toStats(List<PokemonStat> statsList) {
        return statsList.stream()
                .map(stat -> new PokemonStatDocument(stat.getName(), stat.getValue()))
                .toList();
    }

    private PokemonSpecialDocument toSpecial(PokemonSpecial special) {
        return new PokemonSpecialDocument(
                special.getIsSpecial(),
                special.getLegendary(),
                special.getMythical(),
                special.getBaby());
    }

    /**
     * When using this in k8s minikube, and we want to access file-api on localhost.
     * then this is not the "true" address - since tunneling is addressing randomly
     * but… it works... in theory…
     * see - "minikube service file-api-service -n pokedex --url"
     * see - https://minikube.sigs.k8s.io/docs/handbook/accessing/
     */
    public PokemonMediaDocument toDocument(PokemonMediaValue request) {
        var src = UriComponentsBuilder.fromUriString(fileApiProperties.getHost())
                .path(fileApiProperties.getUri())
                .build(request.getFileName());
        return new PokemonMediaDocument(
                request.getMediaId(),
                src.toString(),
                request.getFileName()
        );
    }
}
