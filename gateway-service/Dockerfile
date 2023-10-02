FROM  --platform=linux/amd64 maven:latest
WORKDIR /app
COPY . .
RUN mvn clean install
RUN mvn package


FROM  --platform=linux/amd64 openjdk:17
COPY target/gateway-service-0.0.1-SNAPSHOT.jar /app/gateway-service-0.0.1-SNAPSHOT.jar

EXPOSE 9090

CMD java -jar /app/gateway-service-0.0.1-SNAPSHOT.jar