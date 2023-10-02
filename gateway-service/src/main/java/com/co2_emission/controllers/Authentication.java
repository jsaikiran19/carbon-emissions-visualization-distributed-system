package com.co2_emission.controllers;

import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.concurrent.TimeoutException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.reactive.function.client.WebClient;

import com.co2_emission.model.BaseUrl;
import com.co2_emission.model.Login;
import com.co2_emission.model.SignUp;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@RestController
@CrossOrigin(origins = "*")
public class Authentication {

	@Autowired
	public BaseUrl url;
	
	@JsonDeserialize
    @PostMapping(value = "/login")
	public ResponseEntity<Object> login(@RequestBody Login login) throws IOException, TimeoutException, InterruptedException, URISyntaxException {
		WebClient client = WebClient.builder().build();
		ResponseEntity<Object> response = client.post().uri(new URI(url.Registry+"/user/login"))
			    .bodyValue(login)
			    .retrieve()
			    .toEntity(Object.class).block();
		
		return response;

	}
	
	@JsonDeserialize
	@PostMapping(value = "signUp")
	public ResponseEntity<Object> signUp(@RequestBody SignUp signUp) throws IOException, TimeoutException, InterruptedException, URISyntaxException  {
		WebClient client = WebClient.builder().build();
		ResponseEntity<Object> response = client.post().uri(new URI(url.Registry+"/user/login"))
			    .bodyValue(signUp)
			    .retrieve()
			    .toEntity(Object.class).block();
		
		return response;
	}
}
