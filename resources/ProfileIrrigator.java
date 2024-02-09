package {{.BasePackage}}.page.irrigator;

import {{.BasePackage}}.security.AuthenticationInfos;
import {{.BasePackage}}.repository.BaseUserRepository;
import {{.BasePackage}}.entity.BaseUser;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.ui.Model;

import java.util.Optional;

@Service
@RequiredArgsConstructor
public class ProfileIrrigator {

    private final BaseUserRepository baseUserRepository;
    public void irrigateModel(Model model, AuthenticationInfos userInfos){
        Optional<BaseUser> baseUser = baseUserRepository.findById(userInfos.getId());
        if (!baseUser.isPresent()){
            throw new RuntimeException();
        }
        model.addAttribute("_user", baseUser);
    }
}

