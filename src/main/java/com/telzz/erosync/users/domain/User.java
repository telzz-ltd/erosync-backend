package com.telzz.erosync.users.domain;

import java.time.LocalDateTime;
import java.util.regex.Pattern;

import org.springframework.data.annotation.Id;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class User {
    public enum UserStatus {
        ACTIVE,
        INACTIVE,
        SUSPENDED
    }

    public enum UserRole {
        USER,
        ADMIN,
        MODERATOR
    }

    @Id
    private String id;
    private String name;
    private String email;
    private String passwordHash;
    private UserRole role;
    private UserStatus status;
    private LocalDateTime createdAt;
    private LocalDateTime updatedAt;
    private LocalDateTime emailVerifiedAt;

    private static final Pattern NAME_PATTERN = Pattern.compile("^[a-zA-Z]{3,}(?: [a-zA-Z]{3,}){1,2}$");

    private User() {
    }

    public static User create(String id, String name, String email, String passwordHash) {
        id = trim(id);
        name = trim(name);
        email = trim(email);
        passwordHash = trim(passwordHash);

        if (isEmpty(id) || isEmpty(name) || isEmpty(email) || isEmpty(passwordHash)) {
            throw new IllegalArgumentException("all fields must not be empty");
        }

        if (passwordHash.length() < 8) {
            throw new IllegalArgumentException("password must be 8 or more chars");
        }

        validateEmail(email);

        if (!NAME_PATTERN.matcher(name).matches()) {
            throw new IllegalArgumentException("invalid name");
        }

        LocalDateTime now = LocalDateTime.now();

        User user = new User();
        user.id = id;
        user.name = name;
        user.email = email;
        user.passwordHash = passwordHash;
        user.role = UserRole.USER;
        user.status = UserStatus.ACTIVE;
        user.createdAt = now;
        user.updatedAt = now;

        return user;
    }

    public void makeAdmin() {
        this.role = UserRole.ADMIN;
        touch();
    }

    public void deactivate() {
        this.status = UserStatus.INACTIVE;
        touch();
    }

    public void activate() {
        this.status = UserStatus.ACTIVE;
        touch();
    }

    public void suspend() {
        this.status = UserStatus.SUSPENDED;
        touch();
    }

    public void makeModerator() {
        this.role = UserRole.MODERATOR;
        touch();
    }

    public void makeUser() {
        this.role = UserRole.USER;
        touch();
    }

    public void verifyEmail() {
        this.emailVerifiedAt = LocalDateTime.now();
        touch();
    }

    public boolean isEmailVerified() {
        return this.emailVerifiedAt != null;
    }

    private void touch() {
        this.updatedAt = LocalDateTime.now();
    }

    private static String trim(String value) {
        return value == null ? null : value.trim();
    }

    private static boolean isEmpty(String value) {
        return value == null || value.isBlank();
    }

    private static void validateEmail(String email) {
        if (email.length() > 254 || !email.matches("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")) {
            throw new IllegalArgumentException("invalid email");
        }
    }
}