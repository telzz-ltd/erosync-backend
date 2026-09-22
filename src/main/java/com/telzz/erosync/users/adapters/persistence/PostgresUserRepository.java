package com.telzz.erosync.users.adapters.persistence;

import java.util.List;
import java.util.Optional;

import org.springframework.jdbc.core.JdbcTemplate;

import com.telzz.erosync.users.domain.User;
import com.telzz.erosync.users.ports.UserRepository;

import lombok.RequiredArgsConstructor;

@RequiredArgsConstructor
public class PostgresUserRepository implements UserRepository {
    private final JdbcTemplate jdbcTemplate;

    public void create(User user) {
        jdbcTemplate.update("""
                    INSERT into users
                        (id, name, email, password_hash, status, role, created_at, updated_at, email_verified_at)
                    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
                """,
                user.getId(),
                user.getName(),
                user.getEmail(),
                user.getPasswordHash(),
                user.getStatus(),
                user.getRole(),
                user.getCreatedAt(),
                user.getUpdatedAt(),
                user.getEmailVerifiedAt());
    }

    public void update(User user) {
        jdbcTemplate.update("""
                UPDATE users
                SET
                    name=?,
                    email=?,
                    password_hash=?,
                    status=?,
                    role=?,
                    updated_at=?,
                    email_verified_at=?
                WHERE id = ?;
                """,
                user.getName(),
                user.getEmail(),
                user.getPasswordHash(),
                user.getStatus(),
                user.getRole(),
                user.getUpdatedAt(),
                user.getEmailVerifiedAt(),
                user.getId());
    }

    public List<User> find() {
        return List.of();
    }

    public Optional<User> findById(String id) {
        return null;
    }

    public Optional<User> findByEmail(String email) {
        return null;
    }
}
