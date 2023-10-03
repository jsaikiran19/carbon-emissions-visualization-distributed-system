package com.co2_emission.config;
import java.io.IOException;
import java.io.InputStream;
import java.util.Properties;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class KafkaConfig {

    @Value("${config.file}")
    private String configFile;

    @Bean
    public Properties kafkaProperties() throws IOException {
        try (InputStream inputStream = getClass().getClassLoader().getResourceAsStream(configFile)) {
            if (inputStream == null) {
                throw new IOException(configFile + " not found.");
            }
            Properties properties = new Properties();
            properties.load(inputStream);
            return properties;
        }
    }
}
