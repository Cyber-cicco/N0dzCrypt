package {{.BasePackage}}.security.service;

import {{.BasePackage}}.entity.enums.RoleType;
import {{.BasePackage}}.security.AuthenticationInfos;
import {{.BasePackage}}.util.hx.HX;
import io.jsonwebtoken.Claims;
import jakarta.servlet.http.HttpServletResponse;
import lombok.RequiredArgsConstructor;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.Collections;
import java.util.List;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class AuthenticationService {

    private final JwtService jwtService;
    public AuthenticationInfos getAuthInfos(String token){
        Claims claims = jwtService.extractAllClaims(token);
        return AuthenticationInfos.builder()
                .email(jwtService.extractEmail(claims))
                .roles(jwtService.extractRoles(claims))
                .id(jwtService.extractId(claims))
                .token(token)
                .build();
    }

    public AuthenticationInfos getAuthInfos(){
        return AuthenticationInfos.builder()
                .email(SecurityContextHolder.getContext().getAuthentication().getPrincipal().toString())
                .roles(SecurityContextHolder.getContext().getAuthentication().getAuthorities()
                        .stream()
                        .map(Object::toString)
                        .collect(Collectors.toList()))
                .id((Long) SecurityContextHolder.getContext().getAuthentication().getCredentials())
                .build();
    }

    public void mustBeOfRole(List<String> currentRoles, RoleType expectedRole, HttpServletResponse response) {
        if(!currentRoles.contains(expectedRole.getName())) {
            try {
                response.setHeader(HX.RETARGET, "html");
                response.sendRedirect("/login");
            } catch (IOException e){
                throw new RuntimeException();
            }
        }
    }

    public void rolesMustMatchOne(List<String> currentRoles, List<RoleType> acceptedRoles, HttpServletResponse response) {
        if(Collections.disjoint(acceptedRoles
                .stream()
                .map(RoleType::getName)
                .collect(Collectors.toList()), currentRoles)){
            try {
                response.setHeader(HX.RETARGET, "html");
                response.sendRedirect("/login");
            } catch (IOException e){
                throw new RuntimeException();
            }
        }
    }
}
