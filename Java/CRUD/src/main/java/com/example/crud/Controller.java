package com.example.crud;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.File;
import java.io.FileOutputStream;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;
import java.util.List;
import java.util.Objects;
import java.util.UUID;

@RestController
public class Controller {

    private final  Services services;

    public Controller(Services services){
        this.services = services;
    }

    @GetMapping("/")
    public ResponseEntity<List<Todos>>getTodos(){
        List<Todos> todos = services.findAllTodos();
        return  new ResponseEntity<>(todos,HttpStatus.OK);
    }

    @PostMapping("/")
    public ResponseEntity<String>createTodos(@RequestBody  Todos todo) {

        services.create(todo);
        return new ResponseEntity<>("Created",HttpStatus.CREATED);
    }

    @PostMapping("/img/{id}")
    public  ResponseEntity<String>upload(@RequestParam("file") MultipartFile file, @PathVariable UUID id)  {

        services.UploadImg(file, id);
        return new ResponseEntity<>("Upload", HttpStatus.OK);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<HttpStatus>delete(@PathVariable UUID id){
        services.deleted(id);
        return new ResponseEntity<>(HttpStatus.OK);

    }
    @PutMapping("/{id}")
    public ResponseEntity<HttpStatus> update(@RequestBody Todos todo, @PathVariable UUID id){
        services.update(todo,id);
        return  new ResponseEntity<>(HttpStatus.OK);

    }


}
