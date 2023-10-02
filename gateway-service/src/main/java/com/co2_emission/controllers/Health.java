package com.co2_emission.controllers;

import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@CrossOrigin(origins = "*")
public class Health {

    @GetMapping(value = "/environment")
    public String healthCheck(){
        return new String("App running");
    }
}