package {{.BasePackage}}.page.irrigator;

import {{.BasePackage}}.security.AuthenticationInfos;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.ui.Model;

@Service
@RequiredArgsConstructor
public class {{.ClassName}}Irrigator {

    public void irrigateBase{{.ClassName}}(AuthenticationInfos userInfos, Model model) {
    }
}

