package com.telzz.erosync.users.adapters.persistence;

import java.util.List;
import java.util.Optional;

import org.springframework.stereotype.Repository;

import com.telzz.erosync.users.UserRepository;
import com.telzz.erosync.users.domain.User;

@Repository
public class PostgresUserRepository implements UserRepository {
    public void create(User user) {
    }

    public void update(User user) {

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
