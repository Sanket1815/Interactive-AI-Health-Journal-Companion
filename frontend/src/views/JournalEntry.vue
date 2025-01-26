<template>
  <div class="journal-wrapper">
    <div class="journal-inner">
      <div class="card">
        <div class="card-body">
          <h4 class="card-title mb-1">What's on your mind today? 📝</h4>
          <p class="card-text mb-2">Write your thoughts below</p>
          <form @submit.prevent="createJournal" class="journal-form">
            <!-- Textarea Field -->
            <div class="mb-1 textarea-group">
              <textarea
                v-model="content"
                class="form-control custom-textarea"
                id="content"
                placeholder="How are you feeling today?"
                required
              ></textarea>
              <span class="textarea-icon">&#x1F4DD;</span>
            </div>

            <!-- Submit Button -->
            <button type="submit" class="btn btn-primary w-100 custom-btn">
              Submit Journal
            </button>
          </form>

          <!-- Display Analysis -->
          <div v-if="analysis" class="analysis-result mt-3">
            <h5>Analysis Result:</h5>
            <p>{{ analysis }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "JournalEntry",
  data() {
    return {
      content: "",
      analysis: "", // To store analysis result
    };
  },
  methods: {
    async createJournal() {
      const token = localStorage.getItem("jwt");
      if (!token) {
        alert("You are not logged in!");
        this.$router.push("/login");
        return;
      }

      if (!this.content.trim()) {
        alert("Journal content cannot be empty.");
        return;
      }

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

        // Handle response
        alert(data.message); // Message from the backend
        this.analysis = data.analysis; // Set the analysis result
        this.content = ""; // Clear the input after successful submission
      } catch (error) {
        alert("Error creating journal entry. Please try again later.");
        console.error("Response error:", error);
      }
    },
  },
};
</script>

<style scoped>
.journal-wrapper {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background: linear-gradient(135deg, #ff9a9e 0%, #fad0c4 100%);
}

.journal-inner {
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

.custom-textarea {
  padding: 0.75rem;
  border-radius: 5px;
  border: 1px solid #ccc;
  width: calc(100% - 30px);
  height: 150px;
  resize: none;
  transition: border-color 0.3s ease;
}

.custom-textarea:focus {
  border-color: #ff9a9e;
  outline: none;
}

.textarea-group {
  position: relative;
  margin-bottom: 1rem;
}

.textarea-icon {
  position: absolute;
  right: 10px;
  top: 10px;
  font-size: 1.2rem;
  color: #ff9a9e;
}

.custom-btn {
  background-color: #ff9a9e;
  border: none;
  padding: 0.75rem;
  font-size: 1rem;
  border-radius: 5px;
  transition: background-color 0.3s ease;
}

.custom-btn:hover {
  background-color: #fad0c4;
}

.analysis-result {
  margin-top: 1rem;
  padding: 1rem;
  background-color: #f9f9f9;
  border-left: 5px solid #ff9a9e;
  border-radius: 5px;
  box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.05);
}
</style>
