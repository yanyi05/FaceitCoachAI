import { NavLink } from "react-router-dom";
import styles from "./Sidebar.module.css";

function Sidebar() {
  return (
    <aside className={styles.sidebar}>
      <div className={styles.logo}>
        <h2>Faceit Coach AI</h2>
      </div>

      <nav className={styles.menu}>
        <NavLink
          to="/"
          end
          className={({ isActive }) =>
            isActive ? styles.active : styles.link
          }
        >
          Dashboard
        </NavLink>

        <NavLink
          to="/workspace"
          className={({ isActive }) =>
            isActive ? styles.active : styles.link
          }
        >
          Workspace
        </NavLink>

        <NavLink
          to="/history"
          className={({ isActive }) =>
            isActive ? styles.active : styles.link
          }
        >
          History
        </NavLink>

        <NavLink
          to="/settings"
          className={({ isActive }) =>
            isActive ? styles.active : styles.link
          }
        >
          Settings
        </NavLink>
      </nav>

      <div className={styles.footer}>
        <span>Version 0.1.0</span>
      </div>
    </aside>
  );
}

export default Sidebar;