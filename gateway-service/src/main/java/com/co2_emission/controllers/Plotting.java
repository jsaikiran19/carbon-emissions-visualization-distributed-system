package com.co2_emission.controllers;

import java.io.IOException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import com.co2_emission.components.Consumer;
import com.co2_emission.config.KafkaConfig;

@RestController
@CrossOrigin(value = "*")
public class Plotting {
    
    @Autowired KafkaConfig config;
    @GetMapping(value = "/consume")
    public void consumeMessage() throws IOException{
        Consumer consumer = new Consumer(config.kafkaProperties());
        consumer.consume();
    }
    
}
