import { useNavigate } from "react-router-dom";

import styles from "./Dashboard.module.css";
import { openDemo } from "../../services/demo/demoService";
import Button from "../../components/ui/Button/Button";

function Dashboard() {
  const navigate = useNavigate();

  async function handleSelectDemo() {
    const demo = await openDemo();

    if (!demo) {
      return;
    }

    navigate("/workspace");
  }

  return (
    <div className={styles.dashboard}>
      <h1>Dashboard</h1>

      <p className={styles.subtitle}>
        Welcome back! Analyze your CS2 demos with AI.
      </p>

      <Button onClick={handleSelectDemo}>
        Select Demo
      </Button>
    </div>
  );
}

export default Dashboard;