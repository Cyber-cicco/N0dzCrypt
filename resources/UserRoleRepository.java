package {{.BasePackage}}.repository;

import {{.BasePackage}}.entity.UserRole;
import org.springframework.data.jpa.repository.JpaRepository;

public interface UserRoleRepository extends JpaRepository<UserRole, Long>  {

}

