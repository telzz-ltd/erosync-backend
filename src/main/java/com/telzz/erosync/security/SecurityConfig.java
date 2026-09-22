package com.telzz.erosync.security;

import javax.sql.DataSource;

import org.postgresql.ds.PGPoolingDataSource;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.web.SecurityFilterChain;

import com.telzz.erosync.users.adapters.persistence.PostgresUserRepository;
import com.telzz.erosync.users.ports.UserRepository;

@EnableWebSecurity
@Configuration
public class SecurityConfig {
    @Bean
    public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception {
        return http.build();
    }

    @Bean
    public PasswordEncoder passwordHasher() {
        return new BCryptPasswordEncoder(12);
    }

    @Bean
    public JdbcTemplate jdbcTemplate() {
        DataSource dataSource = new PGPoolingDataSource
        return new JdbcTemplate();
    }

    @Bean
    public UserRepository userRepository() {
        return new PostgresUserRepository(jdbcTemplate());
    }
}
