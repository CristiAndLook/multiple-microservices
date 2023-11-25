FROM mcr.microsoft.com/openjdk/jdk:17-ubuntu AS base

EXPOSE 8080

ADD target/java-microservice.jar java-microservice.jar

ENTRYPOINT [ "java", "-jar", "java-microservice.jar" ]

# docker build -t java-microservice .
# docker run -it --rm -p 8080:8080 --name my-java-microservice java-microservice