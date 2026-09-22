package com.telzz.erosync.security;

import java.time.Instant;
import java.time.temporal.ChronoUnit;
import java.util.Date;

import javax.crypto.SecretKey;

import org.springframework.stereotype.Component;

import com.telzz.erosync.users.domain.User;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jws;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.security.Keys;

@Component
public class JwtService {
    private SecretKey key = Keys.hmacShaKeyFor("my-insecure-256-bits-jwt-secret-key".getBytes());

    public String generateToken(User user) {
        Jwts
                .builder()
                .subject(user.getId())
                .claim("role", user.getRole().name())
                .expiration(Date.from(Instant.now().plus(1, ChronoUnit.HOURS)))
                .signWith(key)
                .compact();
        return "";
    }

    private Claims validateToken(String tokenString) {
        Jws<Claims> claims = Jwts.parser()
                .verifyWith(key)
                .build()
                .parseSignedClaims(tokenString);

        return claims.getPayload();
    }

    public String getUserId(String token) {
        return validateToken(token).getSubject();
    }

    public User.UserRole getUserRole(String token) {
        return validateToken(token).get("role", User.UserRole.class);
    }
}