services:
  app:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: {{ .ProjectName }}_app
    restart: unless-stopped
    env_file:
      - .env
    ports:
      - "{{ .Port }}:{{ .Port }}"
    depends_on:
      - mysql

  mysql:
    image: mysql:8.0
    container_name: {{ .ProjectName }}_mysql
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: ${DB_PASSWORD}
      MYSQL_DATABASE: ${DB_NAME}
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

volumes:
  mysql_data:
