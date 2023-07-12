package com.example.crud;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.stereotype.Component;

import java.io.File;
import java.io.IOException;
import java.util.List;
import java.util.Optional;


@Component
public class FileManager {

    private final String filePath = "db.json";

    private final ObjectMapper objectMapper = new ObjectMapper();
    private final File fileJson = new File(this.filePath);

    public Optional<List<Todos>> readFile() throws IOException {
        List<Todos> todosList = objectMapper.readValue(fileJson, new TypeReference<>() {
        });
        return Optional.ofNullable(todosList);
    }
    public void writeFile(Todos todo) throws IOException {
        List<Todos> todos =  this.readFile().orElseThrow(()-> new  RuntimeException("Error Reading File"));
        todos.add(todo);
        objectMapper.writeValue(fileJson,todos);


    }
}
