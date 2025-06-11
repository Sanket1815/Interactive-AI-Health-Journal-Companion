<template>
  <div class="home-container">
    <div class="background-animation" ref="vantaRef"></div>
    
    <nav class="navbar">
      <div class="nav-content">
        <div class="logo">
          <Brain class="logo-icon" />
          <span class="logo-text">MindJournal</span>
        </div>
        <div class="nav-actions">
          <router-link to="/login" class="btn btn-secondary">Sign In</router-link>
          <router-link to="/register" class="btn btn-primary">Get Started</router-link>
        </div>
      </div>
    </nav>

    <main class="hero-section">
      <div class="hero-content">
        <div class="hero-text">
          <h1 class="hero-title">
            Your AI-Powered
            <span class="gradient-text">Mental Health</span>
            Companion
          </h1>
          <p class="hero-description">
            Transform your thoughts into insights with our intelligent journaling platform. 
            Track your emotions, discover patterns, and receive personalized guidance on your wellness journey.
          </p>
          
          <div class="hero-features">
            <div class="feature-item">
              <Sparkles class="feature-icon" />
              <span>AI-Powered Analysis</span>
            </div>
            <div class="feature-item">
              <Shield class="feature-icon" />
              <span>Private & Secure</span>
            </div>
            <div class="feature-item">
              <TrendingUp class="feature-icon" />
              <span>Track Progress</span>
            </div>
          </div>

          <div class="hero-actions">
            <router-link to="/register" class="btn btn-primary btn-large">
              Start Your Journey
              <ArrowRight class="btn-icon" />
            </router-link>
            <router-link to="/login" class="btn btn-secondary btn-large">
              Welcome Back
            </router-link>
          </div>
        </div>

        <div class="hero-visual">
          <div class="floating-card">
            <div class="card-header">
              <div class="card-dots">
                <span class="dot dot-red"></span>
                <span class="dot dot-yellow"></span>
                <span class="dot dot-green"></span>
              </div>
              <span class="card-title">Today's Entry</span>
            </div>
            <div class="card-content">
              <p class="journal-text">
                "Today I felt grateful for the small moments of joy..."
              </p>
              <div class="sentiment-indicator">
                <div class="sentiment-bar">
                  <div class="sentiment-fill positive"></div>
                </div>
                <span class="sentiment-label">Positive • 85%</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <section class="features-section">
      <div class="features-container">
        <h2 class="section-title">Why Choose MindJournal?</h2>
        <div class="features-grid">
          <div class="feature-card">
            <div class="feature-icon-wrapper">
              <MessageCircle class="feature-card-icon" />
            </div>
            <h3 class="feature-title">Intelligent Conversations</h3>
            <p class="feature-description">
              Engage with our empathetic AI assistant that understands your emotions and provides thoughtful responses.
            </p>
          </div>
          
          <div class="feature-card">
            <div class="feature-icon-wrapper">
              <BarChart3 class="feature-card-icon" />
            </div>
            <h3 class="feature-title">Mood Analytics</h3>
            <p class="feature-description">
              Visualize your emotional patterns over time with detailed sentiment analysis and mood tracking.
            </p>
          </div>
          
          <div class="feature-card">
            <div class="feature-icon-wrapper">
              <Lock class="feature-card-icon" />
            </div>
            <h3 class="feature-title">Complete Privacy</h3>
            <p class="feature-description">
              Your thoughts are sacred. We use end-to-end encryption to keep your journal entries completely private.
            </p>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { Brain, Sparkles, Shield, TrendingUp, ArrowRight, MessageCircle, BarChart3, Lock } from 'lucide-vue-next'
import * as THREE from 'three'
import VANTA from 'vanta'

export default {
  name: "UserHome",
  components: {
    Brain,
    Sparkles,
    Shield,
    TrendingUp,
    ArrowRight,
    MessageCircle,
    BarChart3,
    Lock
  },
  mounted() {
    this.initVanta()
  },
  beforeUnmount() {
    if (this.vantaEffect) {
      this.vantaEffect.destroy()
    }
  },
  methods: {
    initVanta() {
      this.vantaEffect = VANTA.NET({
        el: this.$refs.vantaRef,
        THREE: THREE,
        mouseControls: true,
        touchControls: true,
        gyroControls: false,
        minHeight: 200.00,
        minWidth: 200.00,
        scale: 1.00,
        scaleMobile: 1.00,
        color: 0x667eea,
        backgroundColor: 0xf7fafc,
        points: 8.00,
        maxDistance: 20.00,
        spacing: 16.00
      })
    }
  }
}
</script>

<style scoped>
.home-container {
  min-height: 100vh;
  position: relative;
  overflow-x: hidden;
}

.background-animation {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
}

.navbar {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 100;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid rgba(226, 232, 240, 0.5);
}

.nav-content {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.logo-icon {
  width: 2rem;
  height: 2rem;
  color: #667eea;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: 700;
  color: #2d3748;
}

.nav-actions {
  display: flex;
  gap: 1rem;
}

.hero-section {
  padding: 8rem 2rem 4rem;
  max-width: 1200px;
  margin: 0 auto;
}

.hero-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 4rem;
  align-items: center;
}

.hero-title {
  font-size: 3.5rem;
  font-weight: 800;
  line-height: 1.1;
  margin-bottom: 1.5rem;
  color: #2d3748;
}

.gradient-text {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.hero-description {
  font-size: 1.25rem;
  color: #4a5568;
  margin-bottom: 2rem;
  line-height: 1.6;
}

.hero-features {
  display: flex;
  gap: 2rem;
  margin-bottom: 3rem;
}

.feature-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  color: #667eea;
  font-weight: 500;
}

.feature-icon {
  width: 1.25rem;
  height: 1.25rem;
}

.hero-actions {
  display: flex;
  gap: 1rem;
}

.btn-large {
  padding: 1rem 2rem;
  font-size: 1rem;
}

.btn-icon {
  width: 1rem;
  height: 1rem;
}

.hero-visual {
  display: flex;
  justify-content: center;
  align-items: center;
}

.floating-card {
  background: white;
  border-radius: 1rem;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  overflow: hidden;
  width: 100%;
  max-width: 400px;
  animation: float 6s ease-in-out infinite;
}

@keyframes float {
  0%, 100% { transform: translateY(0px); }
  50% { transform: translateY(-20px); }
}

.card-header {
  padding: 1rem 1.5rem;
  background: #f7fafc;
  border-bottom: 1px solid #e2e8f0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-dots {
  display: flex;
  gap: 0.5rem;
}

.dot {
  width: 0.75rem;
  height: 0.75rem;
  border-radius: 50%;
}

.dot-red { background: #f56565; }
.dot-yellow { background: #ed8936; }
.dot-green { background: #48bb78; }

.card-title {
  font-weight: 600;
  color: #4a5568;
}

.card-content {
  padding: 1.5rem;
}

.journal-text {
  color: #2d3748;
  margin-bottom: 1rem;
  font-style: italic;
}

.sentiment-indicator {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.sentiment-bar {
  flex: 1;
  height: 0.5rem;
  background: #e2e8f0;
  border-radius: 0.25rem;
  overflow: hidden;
}

.sentiment-fill {
  height: 100%;
  border-radius: 0.25rem;
  transition: width 0.3s ease;
}

.sentiment-fill.positive {
  width: 85%;
  background: linear-gradient(90deg, #48bb78, #38a169);
}

.sentiment-label {
  font-size: 0.875rem;
  font-weight: 500;
  color: #48bb78;
}

.features-section {
  padding: 6rem 2rem;
  background: white;
}

.features-container {
  max-width: 1200px;
  margin: 0 auto;
}

.section-title {
  font-size: 2.5rem;
  font-weight: 700;
  text-align: center;
  margin-bottom: 3rem;
  color: #2d3748;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.feature-card {
  text-align: center;
  padding: 2rem;
  border-radius: 1rem;
  transition: transform 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-5px);
}

.feature-icon-wrapper {
  display: inline-flex;
  padding: 1rem;
  background: linear-gradient(135deg, #667eea, #764ba2);
  border-radius: 1rem;
  margin-bottom: 1.5rem;
}

.feature-card-icon {
  width: 2rem;
  height: 2rem;
  color: white;
}

.feature-title {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1rem;
  color: #2d3748;
}

.feature-description {
  color: #4a5568;
  line-height: 1.6;
}

@media (max-width: 768px) {
  .hero-content {
    grid-template-columns: 1fr;
    gap: 2rem;
    text-align: center;
  }
  
  .hero-title {
    font-size: 2.5rem;
  }
  
  .hero-features {
    justify-content: center;
    flex-wrap: wrap;
  }
  
  .hero-actions {
    flex-direction: column;
    align-items: center;
  }
  
  .nav-content {
    padding: 1rem;
  }
  
  .nav-actions {
    gap: 0.5rem;
  }
}
</style>