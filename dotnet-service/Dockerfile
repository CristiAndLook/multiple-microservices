# Create example dotnet service
# Run dotnet new webapi -n dotnet-service --no-https

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /app
COPY dotnet-service.csproj .
RUN dotnet restore
COPY . .

RUN dotnet build "dotnet-service.csproj" -c Release -o /app/build

RUN dotnet publish -c Release -o /app

FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS runtime
WORKDIR /app
COPY --from=build /app .
ENTRYPOINT ["dotnet", "dotnet-service.dll"]

# docker run -it --rm -p 8080:80 --name service-dotnet dotnet-microservice
