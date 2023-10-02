// package com.co2_emission.services;
// import java.io.*;
// import java.nio.file.*;
// import java.util.*;

// import org.jvnet.hk2.annotations.Service;

// @Service
// public class Config {
//     public static Properties loadConfig(final String configFile) throws IOException {
//         if (!Files.exists(Paths.get(configFile))) {
//             throw new IOException(configFile + " not found.");
//         }
//         final Properties cfg = new Properties();
//         try (InputStream inputStream = new FileInputStream(configFile)) {
//             cfg.load(inputStream);
//         }
//         return cfg;
//     }
//     final Properties props = loadConfig("client.properties");

// }