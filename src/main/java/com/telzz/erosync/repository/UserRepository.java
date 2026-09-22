package com.telzz.erosync.repository;

import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Component;

import com.telzz.erosync.users.domain.User;

import lombok.RequiredArgsConstructor;

@Component
@RequiredArgsConstructor
public class UserRepository {
    private final JdbcTemplate jdbcTemplate;

    public void Save(User user) {
        jdbcTemplate.execute(
                """
                        INSERT into users (id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at)
                        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
                        ON CONFLICT (id) DO UPDATE SET
                            name=EXCLUDED.name,
                            email=EXCLUDED.email,
                            password_hash=EXCLUDED.password_hash,
                            status=EXCLUDED.status,
                            role=EXCLUDED.role,
                            created_at=EXCLUDED.created_at,
                            updated_at=EXCLUDED.updated_at,
                            email_verified_at=EXCLUDED.email_verified_at
                        ;
                        """);
    }
}
