package {{.BasePackage}}.page;


import {{.BasePackage}}.page.irrigator.LayoutIrrigator;
import {{.BasePackage}}.security.service.AuthenticationService;
import {{.BasePackage}}.security.AuthenticationInfos;

import jakarta.servlet.http.HttpServletResponse;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;
import lombok.RequiredArgsConstructor;

@Controller
@RequiredArgsConstructor
@RequestMapping("{{.PageName}}")
public class {{.ClassName}}Controller {

    private final LayoutIrrigator layoutIrrigator;
    private final AuthenticationService authenticationService;
    private final {{.ClassName}}Irrigator {{.PageName}}Irrigator;

    @GetMapping
    public String get{{.ClassName}}(Model model, HttpServletResponse response) {
        AuthenticationInfos userInfos = authenticationService.getAuthInfos();
        {{.PageName}}Irrigator.irrigateBase{{.ClassName}}(userInfos, model);
        layoutIrrigator.irrigateBaseLayout(model, Routes.ADR_{{.UpperClassName}});
        return Routes.ADR_BASE_LAYOUT;
    }
    @GetMapping("/partial")
    public String get{{.ClassName}}Partial(Model model, HttpServletResponse response) {
        AuthenticationInfos userInfos = authenticationService.getAuthInfos();
        {{.PageName}}Irrigator.irrigateBase{{.ClassName}}(userInfos, model);
        return Routes.ADR_{{.UpperClassName}};
    }
}
