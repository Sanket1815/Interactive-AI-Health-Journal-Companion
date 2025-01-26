<template>
  <div class="home-wrapper">
    <div class="home-inner">
      <div class="card">
        <div class="card-body text-center">
          <h1 class="card-title mb-3">Welcome to Health Sentiment 🌟</h1>
          <p class="card-text mb-4">
            A mental health journal app that helps you log your daily thoughts
            and emotions. Track your well-being and receive personalized
            insights.
          </p>

          <!-- Navigation Buttons -->
          <div class="btn-group">
            <router-link to="/login" class="btn custom-btn">Login</router-link>
            <router-link to="/register" class="btn custom-btn"
              >Register</router-link
            >
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "UserHome",
};
</script>

<style scoped>
.home-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(135deg, #8e2de2 0%, #4a00e0 100%);
}

.home-inner {
  max-width: 600px;
  width: 100%;
}

.card {
  border-radius: 10px;
  overflow: hidden;
}

.card-body {
  padding: 2rem;
  background-color: #ffffff;
  box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.1);
}

.card-title {
  font-size: 2rem;
  font-weight: bold;
  color: #4a00e0;
}

.card-text {
  font-size: 1.1rem;
  color: #555555;
}

.btn-group {
  margin-top: 2rem;
  display: flex;
  justify-content: center;
  gap: 1rem;
}

.custom-btn {
  background-color: #4a00e0;
  border: none;
  padding: 0.75rem 1.5rem;
  font-size: 1rem;
  border-radius: 5px;
  color: #fff;
  transition: background-color 0.3s ease;
}

.custom-btn:hover {
  background-color: #8e2de2;
}
</style>
