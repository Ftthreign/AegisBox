<script lang="ts">
  import { Lock, Key, ArrowRight } from "lucide-svelte";
  import { onMount } from "svelte";
  import gsap from "gsap";

  let heroContainer: HTMLDivElement;
  let textBlock: HTMLDivElement;
  let card: HTMLDivElement;

  let displayedPassword = "";
  const finalPassword = "/`+\\hyZ'Zmo,'*=F";
  const chars =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{};:,.<>?";

  function animatePassword() {
    setInterval(() => {
      displayedPassword = Array.from(
        { length: finalPassword.length },
        () => chars[Math.floor(Math.random() * chars.length)]
      ).join("");
    }, 90);
  }

  onMount(() => {
    gsap.from(textBlock, {
      opacity: 0,
      y: 40,
      duration: 4,
      ease: "power3.out"
    });

    gsap.from(card, {
      opacity: 0,
      y: 50,
      duration: 4,
      delay: 0.3,
      ease: "power3.out"
    });

    animatePassword();
  });
</script>

<section class="relative pt-32 pb-24 px-6 overflow-hidden">

  <div class="pointer-events-none absolute inset-0 opacity-[0.09]"
    style="background-image: linear-gradient(to right, rgb(148 163 184) 1px, transparent 1px),
           linear-gradient(to bottom, rgb(148 163 184) 1px, transparent 1px);
           background-size: 40px 40px;">
  </div>
  <div class="absolute bottom-20 right-1/4 w-96 h-96 bg-cyan-500/10 rounded-full blur-3xl"></div>

  <div class="max-w-6xl mx-auto" bind:this={heroContainer}>
    
    <!-- TEXT BLOCK -->
    <div class="text-center mb-16" bind:this={textBlock}>
      <div class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-gray-100 border border-gray-200 mb-8">
        <span class="w-2 h-2 rounded-full bg-blue-600"></span>
        <span class="text-sm font-medium text-gray-700">Open Source Password Manager</span>
      </div>

      <h1 class="text-5xl md:text-6xl lg:text-7xl font-bold mb-6 tracking-tight text-gray-900">
        Your Digital <br />
        <span class="text-blue-600">Security Vault</span>
      </h1>

      <p class="text-lg text-gray-600 max-w-2xl mx-auto mb-10 leading-relaxed">
        A local-first password manager with military-grade encryption. Zero-knowledge, open source, and built for privacy.
      </p>

      <div class="flex flex-col sm:flex-row gap-4 justify-center items-center mb-20">
        <button class="px-8 py-3 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium transition-colors flex items-center gap-2">
          Get Started Free
          <ArrowRight class="w-4 h-4" />
        </button>
        <button class="px-8 py-3 bg-gray-100 hover:bg-gray-200 text-gray-900 rounded-lg font-medium transition-colors">
          View on GitHub
        </button>
      </div>
    </div>

    <!-- CARD -->
    <div class="max-w-3xl mx-auto" bind:this={card}>
      <div class="bg-gray-50 rounded-xl border border-gray-200 p-6 shadow-sm">
        
        <div class="flex items-start justify-between mb-5">
          <div class="flex items-center gap-3">
            <div class="w-10 h-10 rounded-lg bg-blue-600 flex items-center justify-center">
              <Lock class="w-5 h-5 text-white" />
            </div>
            <div>
              <h3 class="font-semibold text-gray-900">GitHub</h3>
              <p class="text-sm text-gray-500">@your_username</p>
            </div>
          </div>

          <div class="flex gap-1.5">
            <button class="p-2 hover:bg-gray-200 rounded-lg transition-colors">
              <Key class="w-4 h-4 text-gray-400" />
            </button>
          </div>
        </div>

        <div class="space-y-3">
          <div>
            <!-- svelte-ignore a11y_label_has_associated_control -->
            <label class="text-xs text-gray-500 mb-2 block font-medium">PASSWORD</label>

            <div class="flex items-center gap-2 bg-white rounded-lg px-4 py-3 border border-gray-200 font-mono text-sm text-gray-700">
              <span class="transition-all duration-75">
                {displayedPassword}
              </span>

              <div class="ml-auto flex gap-1.5">
                <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
                <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
                <div class="w-1.5 h-1.5 rounded-full bg-green-500"></div>
              </div>
            </div>
          </div>

          <div class="flex items-center justify-between text-xs">
            <span class="text-gray-500">Last updated: 2 days ago</span>
            <span class="px-3 py-1 bg-green-50 text-green-700 rounded-full font-medium border border-green-200">
              Strong
            </span>
          </div>
        </div>

      </div>
    </div>

  </div>
</section>
