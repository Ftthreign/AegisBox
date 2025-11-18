<script lang="ts">
  import { Shield, Database, CheckCircle, ArrowRight, Lock } from "lucide-svelte";
  import { onMount } from "svelte";

  let container: HTMLDivElement;
  let leftBlock: HTMLDivElement;
  let rightCard: HTMLDivElement;
  let items: HTMLDivElement[] = [];

  const securityFeatures: string[] = [
    "Per-user encryption with Argon2id",
    "Multi-Factor Authentication (TOTP)",
    "Auto-lock on idle timeout",
    "Brute-force protection",
    "Secure session management",
    "CSRF protection built-in"
  ];

  onMount(async () => {
    const gsap = (await import("gsap")).default;
    const ScrollTrigger = (await import("gsap/ScrollTrigger")).default;

    gsap.registerPlugin(ScrollTrigger);

    gsap.from(container, {
      opacity: 0,
      y: 40,
      duration: 1,
      ease: "power3.out",
      scrollTrigger: {
        trigger: container,
        start: "top 80%"
      }
    });

    gsap.from(leftBlock, {
      opacity: 0,
      x: -40,
      duration: 0.9,
      delay: 0.1,
      ease: "power3.out",
      scrollTrigger: {
        trigger: container,
        start: "top 80%"
      }
    });

    gsap.from(rightCard, {
      opacity: 0,
      x: 40,
      duration: 0.9,
      delay: 0.2,
      ease: "power3.out",
      scrollTrigger: {
        trigger: container,
        start: "top 80%"
      }
    });

    gsap.from(items, {
      opacity: 0,
      x: -20,
      stagger: 0.12,
      duration: 0.7,
      delay: 0.4,
      ease: "power2.out",
      scrollTrigger: {
        trigger: container,
        start: "top 80%"
      }
    });

    gsap.to(container, {
  opacity: 0,
  y: 40,
  duration: 0.8,
  ease: "power2.out",
  scrollTrigger: {
    trigger: container,
    start: "bottom top",  
    toggleActions: "play none none reverse"
  }
});
  });
  
</script>

<section id="security" class="py-24 px-6">
  <div
    class="max-w-6xl mx-auto grid lg:grid-cols-2 gap-12 items-center"
    bind:this={container}
  >
    <div bind:this={leftBlock}>
      <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-green-50 border border-green-200 mb-6">
        <Shield class="w-4 h-4 text-green-600" />
        <span class="text-sm font-medium text-green-700">Security First</span>
      </div>

      <h2 class="text-4xl font-bold mb-6 text-gray-900">
        Bank-Level Security
        <span class="block text-green-600">for Your Data</span>
      </h2>

      <p class="text-gray-600 mb-10 leading-relaxed">
        Built with the highest security standards.
      </p>

      <div class="space-y-3">
        {#each securityFeatures as sf, i}
          <div class="flex items-start gap-3" bind:this={items[i]}>
            <CheckCircle class="w-5 h-5 text-green-600 mt-0.5" />
            <span class="text-gray-700">{sf}</span>
          </div>
        {/each}
      </div>
    </div>

    <div bind:this={rightCard}>
      <div class="bg-gray-50 rounded-xl border border-gray-200 p-6 shadow-sm">

        <div class="flex items-center gap-2.5 mb-6">
          <Database class="w-5 h-5 text-green-600" />
          <h3 class="font-semibold text-gray-900">Encryption Flow</h3>
        </div>

        <div class="space-y-3">
          <div class="flex items-center justify-between p-4 bg-white rounded-lg border border-gray-200">
            <span class="text-gray-700">Your Password</span>
            <ArrowRight class="w-4 h-4 text-gray-400" />
          </div>

          <div class="flex items-center justify-between p-4 bg-white rounded-lg border border-gray-200">
            <span class="text-gray-700">Argon2id Derivation</span>
            <ArrowRight class="w-4 h-4 text-gray-400" />
          </div>

          <div class="flex items-center justify-between p-4 bg-white rounded-lg border border-gray-200">
            <span class="text-gray-700">AES-256-GCM</span>
            <ArrowRight class="w-4 h-4 text-gray-400" />
          </div>

          <div class="flex items-center justify-between p-4 bg-green-50 rounded-lg border border-green-200">
            <span class="font-semibold text-green-700">Secure Storage</span>
            <Lock class="w-4 h-4 text-green-600" />
          </div>
        </div>

        <div class="mt-6 p-4 bg-gray-100 rounded-lg">
          <p class="text-sm text-gray-700">
            <span class="font-semibold text-gray-900">Zero-knowledge:</span>
            Server never sees your plaintext passwords.
          </p>
        </div>

      </div>
    </div>

  </div>
</section>
