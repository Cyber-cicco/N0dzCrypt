package {{.BasePackage}}.entity;

import jakarta.persistence.*;
import lombok.*;

import java.util.List;

/**
 * Représente une utilisateur de l'application
 *
 * @author N0dzCrypt
 */
@Getter
@Setter
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
public class BaseUser {

	@Id
	@GeneratedValue(strategy = GenerationType.IDENTITY)
	private Long id;

	@ManyToMany
	@JoinTable(name = "user_role",
        joinColumns = @JoinColumn(name = "user_id", referencedColumnName = "id"),
        inverseJoinColumns = @JoinColumn(name = "role_id", referencedColumnName = "id")
    )
	private List<UserRole> userRoles;

	@Column(nullable = false, unique = true, length = 30)
	private String email;

	@Column(nullable = false, length = 80)
	private String password;
}
