package com.example.crud;

import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

@Service
public class Services {

    private List<Todos> todos = new ArrayList<>();


    public List<Todos> findAllTodos(){
        return  todos;
    }

    public  void create(Todos todo){
        todo.setId(UUID.randomUUID());
        todos.add(todo);
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
