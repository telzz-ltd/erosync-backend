package com.telzz.erosync.users;

import java.util.List;
import java.util.Optional;

import com.telzz.erosync.users.domain.User;

import jakarta.annotation.Nonnull;

public interface UserRepository {
    List<User> find();

    Optional<User> findById(String id);

    Optional<User> findByEmail(String email);

    void create(@Nonnull User user);

    void update(@Nonnull User user);
}
