<template>
  <div class="journal-container">
    <div class="journal-background" ref="vantaRef"></div>
    
    <nav class="navbar">
      <div class="nav-content">
        <div class="logo">
          <Brain class="logo-icon" />
          <span class="logo-text">MindJournal</span>
        </div>
        <button @click="logout" class="btn btn-secondary">
          <LogOut class="btn-icon" />
          Sign Out
        </button>
      </div>
    </nav>

    <main class="journal-main">
      <div class="journal-content">
        <div class="journal-header">
          <h1 class="journal-title">
            <PenTool class="title-icon" />
            How are you feeling today?
          </h1>
          <p class="journal-subtitle">
            Share your thoughts and let our AI companion provide insights and support
          </p>
        </div>

        <div class="journal-form-container">
          <form @submit.prevent="createJournal" class="journal-form">
            <div class="form-group">
              <label for="content" class="form-label">
                <Heart class="label-icon" />
                Your thoughts and feelings
              </label>
              <div class="textarea-container">
                <textarea
                  v-model="content"
                  id="content"
                  class="journal-textarea"
                  placeholder="What's on your mind? Share your thoughts, feelings, experiences, or anything you'd like to reflect on..."
                  required
                  :disabled="loading"
                  rows="8"
                ></textarea>
                <div class="character-count">
                  {{ content.length }} characters
                </div>
              </div>
            </div>

            <button 
              type="submit" 
              class="btn btn-primary btn-large btn-full"
              :disabled="loading || !content.trim()"
            >
              <Loader2 v-if="loading" class="btn-icon animate-spin" />
              <Sparkles v-else class="btn-icon" />
              {{ loading ? 'Analyzing...' : 'Share & Analyze' }}
            </button>
          </form>

          <transition name="fade" mode="out-in">
            <div v-if="analysis" class="analysis-container">
              <div class="analysis-header">
                <div class="analysis-icon-wrapper">
                  <MessageCircle class="analysis-icon" />
                </div>
                <h3 class="analysis-title">AI Insights</h3>
              </div>
              
              <div class="analysis-content">
                <div class="analysis-text">
                  {{ analysis }}
                </div>
                
                <div class="analysis-actions">
                  <button @click="clearAnalysis" class="btn btn-secondary btn-small">
                    <X class="btn-icon" />
                    Clear
                  </button>
                  <button @click="saveEntry" class="btn btn-primary btn-small">
                    <Save class="btn-icon" />
                    Save Entry
                  </button>
                </div>
              </div>
            </div>
          </transition>
        </div>

        <div class="journal-tips">
          <h4 class="tips-title">
            <Lightbulb class="tips-icon" />
            Journaling Tips
          </h4>
          <div class="tips-grid">
            <div class="tip-item">
              <div class="tip-icon">🎯</div>
              <span>Be honest and authentic with your feelings</span>
            </div>
            <div class="tip-item">
              <div class="tip-icon">🌱</div>
              <span>Focus on growth and self-reflection</span>
            </div>
            <div class="tip-item">
              <div class="tip-icon">💭</div>
              <span>Write about both challenges and victories</span>
            </div>
            <div class="tip-item">
              <div class="tip-icon">🔒</div>
              <span>Your entries are private and secure</span>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { 
  Brain, 
  LogOut, 
  PenTool, 
  Heart, 
  Sparkles, 
  MessageCircle, 
  X, 
  Save, 
  Lightbulb,
  Loader2
} from 'lucide-vue-next'
import * as THREE from 'three'
import VANTA from 'vanta'

export default {
  name: "JournalEntry",
  components: {
    Brain,
    LogOut,
    PenTool,
    Heart,
    Sparkles,
    MessageCircle,
    X,
    Save,
    Lightbulb,
    Loader2
  },
  data() {
    return {
      content: "",
      analysis: "",
      loading: false,
      saved: false
    };
  },
  mounted() {
    this.initVanta()
    this.checkAuth()
  },
  beforeUnmount() {
    if (this.vantaEffect) {
      this.vantaEffect.destroy()
    }
  },
  methods: {
    initVanta() {
      this.vantaEffect = VANTA.BIRDS({
        el: this.$refs.vantaRef,
        THREE: THREE,
        mouseControls: true,
        touchControls: true,
        gyroControls: false,
        minHeight: 200.00,
        minWidth: 200.00,
        scale: 1.00,
        scaleMobile: 1.00,
        backgroundColor: 0xff9a9e,
        color1: 0xfad0c4,
        color2: 0xff9a9e,
        colorMode: "variance",
        birdSize: 1.20,
        wingSpan: 25.00,
        speedLimit: 4.00,
        separation: 20.00,
        alignment: 20.00,
        cohesion: 20.00,
        quantity: 3.00
      })
    },
    checkAuth() {
      const token = localStorage.getItem("jwt");
      if (!token) {
        this.$router.push("/login");
      }
    },
    async createJournal() {
      if (this.loading || !this.content.trim()) return;

      const token = localStorage.getItem("jwt");
      if (!token) {
        alert("You are not logged in!");
        this.$router.push("/login");
        return;
      }

      this.loading = true;
      this.analysis = "";

      try {
        const response = await fetch("http://localhost:8080/journal", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
          body: JSON.stringify({
            content: this.content,
          }),
        });

        if (response.status === 401) {
          alert("Your session has expired. Please log in again.");
          localStorage.removeItem("jwt");
          this.$router.push("/login");
          return;
        }

        if (!response.ok) {
          const errorMessage = await response.text();
          throw new Error(`Failed to create journal entry: ${errorMessage}`);
        }

        const data = await response.json();
        this.analysis = data.analysis;
        this.saved = false;

      } catch (error) {
        alert("Error analyzing your entry. Please try again later.");
        console.error("Journal creation error:", error);
      } finally {
        this.loading = false;
      }
    },
    clearAnalysis() {
      this.analysis = "";
      this.content = "";
      this.saved = false;
    },
    saveEntry() {
      // In a real app, you might want to save to a separate "saved entries" collection
      // For now, we'll just show a success message
      this.saved = true;
      setTimeout(() => {
        alert("Entry saved successfully!");
        this.clearAnalysis();
      }, 500);
    },
    logout() {
      localStorage.removeItem("jwt");
      this.$router.push("/");
    }
  },
};
</script>

<style scoped>
.journal-container {
  min-height: 100vh;
  position: relative;
}

.journal-background {
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
  color: #ff9a9e;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: 700;
  color: #2d3748;
}

.btn-icon {
  width: 1rem;
  height: 1rem;
}

.journal-main {
  padding: 6rem 2rem 2rem;
  max-width: 800px;
  margin: 0 auto;
}

.journal-content {
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 1.5rem;
  padding: 3rem;
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.25);
  border: 1px solid rgba(255, 255, 255, 0.2);
}

.journal-header {
  text-align: center;
  margin-bottom: 3rem;
}

.journal-title {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  font-size: 2.5rem;
  font-weight: 700;
  color: #2d3748;
  margin-bottom: 1rem;
}

.title-icon {
  width: 2.5rem;
  height: 2.5rem;
  color: #ff9a9e;
}

.journal-subtitle {
  font-size: 1.125rem;
  color: #4a5568;
  line-height: 1.6;
}

.journal-form-container {
  margin-bottom: 3rem;
}

.form-group {
  margin-bottom: 2rem;
}

.form-label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1rem;
  font-weight: 600;
  color: #2d3748;
  margin-bottom: 1rem;
}

.label-icon {
  width: 1.25rem;
  height: 1.25rem;
  color: #ff9a9e;
}

.textarea-container {
  position: relative;
}

.journal-textarea {
  width: 100%;
  padding: 1.5rem;
  font-size: 1rem;
  line-height: 1.6;
  color: #2d3748;
  background: white;
  border: 2px solid #e2e8f0;
  border-radius: 1rem;
  resize: vertical;
  min-height: 200px;
  transition: all 0.2s ease;
  font-family: inherit;
}

.journal-textarea:focus {
  outline: none;
  border-color: #ff9a9e;
  box-shadow: 0 0 0 3px rgba(255, 154, 158, 0.1);
}

.journal-textarea::placeholder {
  color: #a0aec0;
}

.character-count {
  position: absolute;
  bottom: 1rem;
  right: 1rem;
  font-size: 0.75rem;
  color: #a0aec0;
  background: rgba(255, 255, 255, 0.9);
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
}

.btn-large {
  padding: 1rem 2rem;
  font-size: 1rem;
  font-weight: 600;
}

.btn-full {
  width: 100%;
  justify-content: center;
}

.btn-small {
  padding: 0.5rem 1rem;
  font-size: 0.875rem;
}

.animate-spin {
  animation: spin 1s linear infinite;
}

@keyframes spin {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}

.analysis-container {
  margin-top: 2rem;
  background: linear-gradient(135deg, #f7fafc 0%, #edf2f7 100%);
  border-radius: 1rem;
  padding: 2rem;
  border: 1px solid #e2e8f0;
}

.analysis-header {
  display: flex;
  align-items: center;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.analysis-icon-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 3rem;
  height: 3rem;
  background: linear-gradient(135deg, #ff9a9e, #fad0c4);
  border-radius: 0.75rem;
}

.analysis-icon {
  width: 1.5rem;
  height: 1.5rem;
  color: white;
}

.analysis-title {
  font-size: 1.5rem;
  font-weight: 600;
  color: #2d3748;
}

.analysis-content {
  margin-bottom: 1.5rem;
}

.analysis-text {
  font-size: 1rem;
  line-height: 1.7;
  color: #4a5568;
  margin-bottom: 1.5rem;
  padding: 1.5rem;
  background: white;
  border-radius: 0.75rem;
  border-left: 4px solid #ff9a9e;
}

.analysis-actions {
  display: flex;
  gap: 1rem;
  justify-content: flex-end;
}

.journal-tips {
  background: linear-gradient(135deg, #fff5f5 0%, #fed7d7 100%);
  border-radius: 1rem;
  padding: 2rem;
  border: 1px solid #feb2b2;
}

.tips-title {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  font-size: 1.25rem;
  font-weight: 600;
  color: #2d3748;
  margin-bottom: 1.5rem;
}

.tips-icon {
  width: 1.5rem;
  height: 1.5rem;
  color: #f56565;
}

.tips-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1rem;
}

.tip-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1rem;
  background: rgba(255, 255, 255, 0.7);
  border-radius: 0.5rem;
  font-size: 0.875rem;
  color: #4a5568;
}

.tip-icon {
  font-size: 1.25rem;
  flex-shrink: 0;
}

.fade-enter-active, .fade-leave-active {
  transition: all 0.3s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

@media (max-width: 768px) {
  .journal-main {
    padding: 6rem 1rem 2rem;
  }
  
  .journal-content {
    padding: 2rem;
  }
  
  .journal-title {
    font-size: 2rem;
    flex-direction: column;
    gap: 0.5rem;
  }
  
  .title-icon {
    width: 2rem;
    height: 2rem;
  }
  
  .analysis-actions {
    flex-direction: column;
  }
  
  .tips-grid {
    grid-template-columns: 1fr;
  }
  
  .nav-content {
    padding: 1rem;
  }
}
</style>