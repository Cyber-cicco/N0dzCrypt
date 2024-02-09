package {{.BasePackage}}.entity;

import {{.BasePackage}}.entity.enums.RoleType;
import jakarta.persistence.*;
import lombok.*;
import org.springframework.security.core.GrantedAuthority;

import java.util.List;

/**
 * @author n0dzCrypt
 */
@AllArgsConstructor
@NoArgsConstructor
@Getter
@Setter
@Builder
@Entity
public class UserRole implements GrantedAuthority {

	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	private Long id;

	@Enumerated
	private RoleType type;

	@ManyToMany
	@JoinTable(name = "user_role",
        joinColumns = @JoinColumn(name = "role_id", referencedColumnName = "id"),
        inverseJoinColumns = @JoinColumn(name = "user_id", referencedColumnName = "id")
    )
	private List<BaseUser> users;

	/** Libellé */
	private String libelle;

	@Override
	public String getAuthority() {
		return type.name();
	}
}

