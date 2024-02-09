package {{.BasePackage}}.repository;

import {{.BasePackage}}.entity.BaseUser;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.Optional;

public interface BaseUserRepository extends JpaRepository<BaseUser, Long>  {

    Optional<BaseUser> findByEmail(String email);
}

