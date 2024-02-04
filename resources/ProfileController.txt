package {{.BasePackage}}.page;

import {{.BasePackage}}.page.irrigator.ProfileIrrigator;
import {{.BasePackage}}.page.irrigator.LayoutIrrigator;
import {{.BasePackage}}.security.AuthenticationInfos;
import {{.BasePackage}}.security.service.AuthenticationService;
import jakarta.servlet.http.HttpServletResponse;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

import java.io.IOException;

@Controller
@RequestMapping
@RequiredArgsConstructor
public class ProfileController {

    private final AuthenticationService authenticationService;
    private final ProfileIrrigator profileIrrigator;
    private final LayoutIrrigator layoutIrrigator;

    @GetMapping({"api", "profile/partial"})
    public String getProfileApi(Model model){
        AuthenticationInfos userInfos = authenticationService.getAuthInfos();
        profileIrrigator.irrigateModel(model, userInfos);
        return Routes.ADR_PROFILE;
    }
    @GetMapping({"/", "", "profile"})
    public String getProfile(Model model, HttpServletResponse response) throws IOException {
        try {
            AuthenticationInfos userInfos = authenticationService.getAuthInfos();
            layoutIrrigator.irrigateBaseLayout(model, Routes.ADR_PROFILE);
            profileIrrigator.irrigateModel(model, userInfos);
            return Routes.ADR_BASE_LAYOUT;
        } catch (Exception e){
            response.sendRedirect("login");
            return Routes.ADR_LOGIN;
        }
    }
}

