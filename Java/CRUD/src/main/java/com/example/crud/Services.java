package com.example.crud;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

@Service
public class Services {

    private List<Todos> todos = new ArrayList<>();

    private final   FileManager fileManager;

    public Services(FileManager fileManager) {
        this.fileManager = fileManager;
    }

    public List<Todos> findAllTodos(){
        try {
            return fileManager.readFile().orElseThrow(()-> new RuntimeException("Error Reading File"));
        }catch (Exception e){
           throw  new RuntimeException(e);
        }
    }

    public Todos findbyId(UUID id){
        try {
            List<Todos> list =   fileManager.readFile().orElseThrow(()-> new RuntimeException("Error "));
            System.out.println(list);
            List<Todos> todos1 =  list.stream().filter((todo)-> todo.getId().equals(id)).toList();
            System.out.println(todos1);
         return todos1.get(0);
        }catch (Exception e){
            throw new RuntimeException(e);
        }

    }
    public  void create(Todos todo){
        todo.setId(UUID.randomUUID());
        try {
            fileManager.writeFile(todo);
        }catch (Exception e){
            throw  new RuntimeException(e);
        }
    }

    public void  deleted(UUID id){

       List<Todos> todosIm = todos.stream().filter((todo)-> !todo.getId().equals(id)).toList();
       this.todos = new ArrayList<>(todosIm);
       System.out.println(todos);

    }

    public void  update(Todos todoRequest,UUID id){
       List<Todos> todoDb =  todos.stream().map(todo -> {
            if(todo.getId().equals(id)){
               todo.setTitle(todoRequest.getTitle());
               todo.setDescription(todoRequest.getDescription());
            }
            return  todo;
        }).toList();

       this.todos = new ArrayList<>(todoDb);


    }


}
