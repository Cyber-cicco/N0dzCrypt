package {{.BasePackage}}.page;

import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import lombok.RequiredArgsConstructor;

@Controller
@RequiredArgsConstructor
@RequestMapping("/home")
public class HomeController {

    @GetMapping
    public String getHome(Model model) {
        model.addAttribute("routerOutlet", Routes.ADR_HOME);
        return Routes.ADR_BASE_LAYOUT;
    }

    @GetMapping("/partial")
    public String getHomePart(Model model) {
        return Routes.ADR_HOME;
    }
}

