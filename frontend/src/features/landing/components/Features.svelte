<script lang="ts">
  import { Lock, Key, Server, Zap } from 'lucide-svelte';
  import { onMount } from 'svelte';

  let activeFeature = 0;

  const features = [
    { icon: Lock, title: "Military-Grade Encryption", description: "AES-256-GCM protects every credential" },
    { icon: Server, title: "Local-First", description: "Your data stays on your device" },
    { icon: Key, title: "Zero-Knowledge", description: "Server never sees your passwords" },
    { icon: Zap, title: "Blazingly Fast", description: "Built with Go and SvelteKit" }
  ];

  // element refs
  let container: HTMLDivElement;
  let cards: HTMLDivElement[] = [];

  onMount(async () => {
    const gsap = (await import("gsap")).default;
    const ScrollTrigger = (await import("gsap/ScrollTrigger")).default;

    gsap.registerPlugin(ScrollTrigger);

    // Fade-in container
    gsap.from(container, {
      opacity: 0,
      y: 30,
      duration: 0.8,
      ease: "power3.out",
      scrollTrigger: {
        trigger: container,
        start: "top 85%"
      }
    });
  });

  onMount(() => {
    const interval = setInterval(() => {
      activeFeature = (activeFeature + 1) % features.length;
    }, 5000);

    return () => clearInterval(interval);
  });
</script>

<section id="features" class="py-24 px-6 bg-gray-50">
  <div class="max-w-6xl mx-auto" bind:this={container}>

    <div class="text-center mb-16">
      <h2 class="text-4xl font-bold mb-3 text-gray-900">Powerful Features</h2>
      <p class="text-gray-600 max-w-xl mx-auto">Everything you need for secure password management</p>
    </div>

    <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-6">
      {#each features as f, i}
        <!-- svelte-ignore a11y_no_static_element_interactions -->
        <div
          class="relative group overflow-hidden p-6 rounded-lg border transition-all cursor-pointer 
            {activeFeature === i
              ? 'bg-white border-blue-300 shadow-lg'
              : 'bg-white border-gray-200 hover:border-gray-300'}"
          on:mouseenter={() => (activeFeature = i)}
          bind:this={cards[i]}
        >

          <div
            class="absolute inset-0 opacity-0 group-hover:opacity-100 transition duration-700 pointer-events-none"
            style="
              background: linear-gradient(
                120deg,
                transparent 0%,
                rgba(255,255,255,0.45) 50%,
                transparent 100%
              );
              transform: translateX(-100%);
            "
          ></div>

          <script>
            onMount(async () => {
              const gsap = (await import("gsap")).default;

              const el = cards[i].querySelector("div");
              cards[i].addEventListener("mouseenter", () => {
                gsap.fromTo(
                  el,
                  { x: "-120%" },
                  { x: "120%", duration: 1, ease: "power2.out" }
                );
              });
            });
          </script>

          <div
            class="w-10 h-10 rounded-lg flex items-center justify-center mb-4
              {activeFeature === i ? 'bg-blue-100 text-blue-600' : 'bg-gray-100 text-gray-600'}">
            <f.icon class="w-5 h-5" />
          </div>

          <h3 class="font-semibold mb-2 text-gray-900">{f.title}</h3>
          <p class="text-sm text-gray-600">{f.description}</p>

        </div>
      {/each}
    </div>

  </div>
</section>
