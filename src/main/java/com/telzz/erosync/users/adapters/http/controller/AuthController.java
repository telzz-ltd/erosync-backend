package com.telzz.erosync.users.adapters.http.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.telzz.erosync.security.JwtService;
import com.telzz.erosync.users.domain.User;
import com.telzz.erosync.users.usecase.CreateUser;
import com.telzz.erosync.users.usecase.ValidateUser;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotEmpty;
import lombok.RequiredArgsConstructor;

@RestController
@RequestMapping("/auth")
@RequiredArgsConstructor
public class AuthController {
    private final CreateUser createUser;
    private final ValidateUser validateUser;
    private final JwtService jwtService;

    public record RegisterRequest(
            @NotEmpty @Min(6) String name,
            @NotEmpty @Email String email,
            @NotEmpty @Min(8) @Max(50) String password) {
    }

    public record AuthResponse(User user, String accessToken) {
    }

    @PostMapping("/register")
    public AuthResponse registerUser(@RequestBody RegisterRequest req) {
        CreateUser.Command cmd = new CreateUser.Command(req.name(), req.email(), req.password());

        User user = createUser.execute(cmd);
        String accessToken = jwtService.generateToken(user);

        return new AuthResponse(user, accessToken);
    }

    public record LoginRequest(
            @NotEmpty @Email String email,
            @NotEmpty @Min(8) @Max(50) String password) {
    }

    @PostMapping("/login")
    public AuthResponse login(@RequestBody LoginRequest req) {
        User user = validateUser.execute(req.email(), req.password());
        String accessToken = jwtService.generateToken(user);

        return new AuthResponse(user, accessToken);
    }
}
