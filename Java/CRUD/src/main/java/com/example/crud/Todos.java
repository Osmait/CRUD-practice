package com.example.crud;


import java.util.UUID;

public class Todos {

    private UUID id;
    private String title;
    private String Description;
    private String img;
    public String getImg() {
        return img;
    }

    public void setImg(String img) {
        this.img = img;
    }



    public UUID getId() {
        return id;
    }

    public void setId(UUID id) {
        this.id = id;
    }

    public String getTitle() {
        return title;
    }

    public void setTitle(String title) {
        this.title = title;
    }

    public String getDescription() {
        return Description;
    }

    public void setDescription(String description) {
        Description = description;
    }

    public Todos(String title, String description) {

        this.title = title;
        this.Description = description;
    }

    @Override
    public String toString() {
        return "Todos{" +
                "id=" + id +
                ", title='" + title + '\'' +
                ", Description='" + Description + '\'' +
                '}';
    }

    public Todos() {

    }
}
