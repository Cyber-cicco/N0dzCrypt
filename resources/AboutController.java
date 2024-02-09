package {{.BasePackage}}.page;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;

@RequiredArgsConstructor
@Controller
@RequestMapping("/about")
public class AboutController {

    @GetMapping
    public String getHomePage(Model model){
        model.addAttribute("routerOutlet", "page/about/about");
        return "layout/base";
    }
    @GetMapping("/partial")
    public String getHomePagePartial(Model model){
        model.addAttribute("routerOutlet", "pages/about");
        return "page/about/about";
    }

}

