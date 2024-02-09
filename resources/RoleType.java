package {{.BasePackage}}.entity.enums;

import lombok.AllArgsConstructor;
import lombok.Getter;

/**
 * @author n0dzCrypt
 */
@Getter
@AllArgsConstructor
public enum RoleType {

	ROLE_USER("User"),
	ROLE_ADMIN("Administrator");

    private String name;

}

