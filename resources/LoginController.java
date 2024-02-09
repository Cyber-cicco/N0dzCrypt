package {{.BasePackage}}.page;

import {{.BasePackage}}.entity.UserRole;
import {{.BasePackage}}.entity.BaseUser;
import {{.BasePackage}}.page.irrigator.LayoutIrrigator;
import {{.BasePackage}}.repository.BaseUserRepository;
import {{.BasePackage}}.security.AuthenticationInfos;
import {{.BasePackage}}.security.service.JwtService;
import {{.BasePackage}}.util.hx.HX;
import jakarta.servlet.http.HttpServletResponse;
import lombok.*;
import org.springframework.http.HttpHeaders;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.Optional;
import java.util.stream.Collectors;

@Controller
@RequiredArgsConstructor
public class LoginController {

    private final BaseUserRepository BaseUserRepository;
    private final JwtService jwtService;
    private final PasswordEncoder passwordEncoder;
    private final LayoutIrrigator layoutIrrigator;

    @GetMapping("/login")
    public String getloginPage(){
        return Routes.ADR_LOGIN;
    }
    @GetMapping("/login/api")
    public void redirectToLogin(HttpServletResponse response) throws IOException {
        response.setHeader(HX.REDIRECT, "/login");
    }

    @Getter
    @Setter
    @AllArgsConstructor
    private static class LoginDto{private String email; private String password;}

    @PostMapping("/login")
    public String login(@ModelAttribute LoginDto loginDto, Model model, HttpServletResponse response){
        Optional<BaseUser> auth = BaseUserRepository.findByEmail(loginDto.email)
                .filter(baseUser -> passwordEncoder.matches(loginDto.password, baseUser.getPassword())) ;
        if(auth.isPresent()){
            response.setHeader(HttpHeaders.SET_COOKIE, jwtService.buildJWTCookie(auth.get()));
            layoutIrrigator.irrigateBaseLayout(model, Routes.ADR_HOME);
            response.setHeader(HX.REDIRECT, "profile");
            return Routes.ADR_BASE_LAYOUT;
        }
        model.addAttribute("error", "Votre adresse e-mail ou mot de passe est invalide");
        return Routes.ADR_FORM_ERROR;
    }
}

