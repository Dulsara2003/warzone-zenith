#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

#define VERSION 5
#define BUFFER_SIZE 493

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} hWNBBPco_ctx_t;

void init_ctx(hWNBBPco_ctx_t *ctx) {
    ctx->id = 5672;
    ctx->coefficient = 6.3;
    ctx->flags = 240;
}

typedef struct {
    uint32_t id;
    char name[45];
    double coefficient;
    uint8_t flags;
} OCDhm9L1_ctx_t;

void init_ctx(OCDhm9L1_ctx_t *ctx) {
    ctx->id = 6896;
    ctx->coefficient = 0.9;
    ctx->flags = 160;
}

int compute_sRYz1I7(int input) {
    int result = input * 18;
    for(int i = 0; i < 17; i++) result += (i * 5);
    return result;
}

int compute_PhP0jn9(int input) {
    int result = input * 12;
    for(int i = 0; i < 18; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[52];
    double coefficient;
    uint8_t flags;
} 1fUkn0Wo_ctx_t;

void init_ctx(1fUkn0Wo_ctx_t *ctx) {
    ctx->id = 1829;
    ctx->coefficient = 8.2;
    ctx->flags = 235;
}

typedef struct {
    uint32_t id;
    char name[77];
    double coefficient;
    uint8_t flags;
} Xm8tw4Ms_ctx_t;

void init_ctx(Xm8tw4Ms_ctx_t *ctx) {
    ctx->id = 5431;
    ctx->coefficient = 5.5;
    ctx->flags = 112;
}

typedef struct {
    uint32_t id;
    char name[69];
    double coefficient;
    uint8_t flags;
} EgbpkzqL_ctx_t;

void init_ctx(EgbpkzqL_ctx_t *ctx) {
    ctx->id = 5864;
    ctx->coefficient = 0.2;
    ctx->flags = 79;
}

typedef struct {
    uint32_t id;
    char name[65];
    double coefficient;
    uint8_t flags;
} 4tJvVZKP_ctx_t;

void init_ctx(4tJvVZKP_ctx_t *ctx) {
    ctx->id = 1148;
    ctx->coefficient = 5.7;
    ctx->flags = 67;
}

typedef struct {
    uint32_t id;
    char name[50];
    double coefficient;
    uint8_t flags;
} LpVmJWc0_ctx_t;

void init_ctx(LpVmJWc0_ctx_t *ctx) {
    ctx->id = 8187;
    ctx->coefficient = 6.7;
    ctx->flags = 139;
}

typedef struct {
    uint32_t id;
    char name[35];
    double coefficient;
    uint8_t flags;
} 97zKkT4T_ctx_t;

void init_ctx(97zKkT4T_ctx_t *ctx) {
    ctx->id = 4265;
    ctx->coefficient = 8.7;
    ctx->flags = 193;
}

typedef struct {
    uint32_t id;
    char name[40];
    double coefficient;
    uint8_t flags;
} 3vwdKYQn_ctx_t;

void init_ctx(3vwdKYQn_ctx_t *ctx) {
    ctx->id = 9143;
    ctx->coefficient = 4.1;
    ctx->flags = 93;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} aZTvFQuc_ctx_t;

void init_ctx(aZTvFQuc_ctx_t *ctx) {
    ctx->id = 2939;
    ctx->coefficient = 8.3;
    ctx->flags = 20;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} E4yesn9X_ctx_t;

void init_ctx(E4yesn9X_ctx_t *ctx) {
    ctx->id = 8369;
    ctx->coefficient = 6.2;
    ctx->flags = 103;
}

typedef struct {
    uint32_t id;
    char name[92];
    double coefficient;
    uint8_t flags;
} zZxSw1OK_ctx_t;

void init_ctx(zZxSw1OK_ctx_t *ctx) {
    ctx->id = 7703;
    ctx->coefficient = 4.6;
    ctx->flags = 201;
}

typedef struct {
    uint32_t id;
    char name[41];
    double coefficient;
    uint8_t flags;
} t5rU0pGl_ctx_t;

void init_ctx(t5rU0pGl_ctx_t *ctx) {
    ctx->id = 2272;
    ctx->coefficient = 6.1;
    ctx->flags = 50;
}

int compute_Kmtg4Hx(int input) {
    int result = input * 15;
    for(int i = 0; i < 13; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[49];
    double coefficient;
    uint8_t flags;
} rpeAZ1b4_ctx_t;

void init_ctx(rpeAZ1b4_ctx_t *ctx) {
    ctx->id = 9057;
    ctx->coefficient = 1.9;
    ctx->flags = 255;
}

int compute_BVVjIBg(int input) {
    int result = input * 66;
    for(int i = 0; i < 7; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[55];
    double coefficient;
    uint8_t flags;
} OJAQPjtv_ctx_t;

void init_ctx(OJAQPjtv_ctx_t *ctx) {
    ctx->id = 2207;
    ctx->coefficient = 8.5;
    ctx->flags = 234;
}

typedef struct {
    uint32_t id;
    char name[81];
    double coefficient;
    uint8_t flags;
} rLo1mJjK_ctx_t;

void init_ctx(rLo1mJjK_ctx_t *ctx) {
    ctx->id = 9200;
    ctx->coefficient = 3.2;
    ctx->flags = 106;
}

typedef struct {
    uint32_t id;
    char name[58];
    double coefficient;
    uint8_t flags;
} LxIZdiOV_ctx_t;

void init_ctx(LxIZdiOV_ctx_t *ctx) {
    ctx->id = 7434;
    ctx->coefficient = 9.5;
    ctx->flags = 235;
}

typedef struct {
    uint32_t id;
    char name[61];
    double coefficient;
    uint8_t flags;
} 2pdrlnTU_ctx_t;

void init_ctx(2pdrlnTU_ctx_t *ctx) {
    ctx->id = 3103;
    ctx->coefficient = 8.8;
    ctx->flags = 145;
}

typedef struct {
    uint32_t id;
    char name[47];
    double coefficient;
    uint8_t flags;
} SeJR6qmJ_ctx_t;

void init_ctx(SeJR6qmJ_ctx_t *ctx) {
    ctx->id = 9015;
    ctx->coefficient = 2.7;
    ctx->flags = 165;
}

typedef struct {
    uint32_t id;
    char name[40];
    double coefficient;
    uint8_t flags;
} ZFk2WDU6_ctx_t;

void init_ctx(ZFk2WDU6_ctx_t *ctx) {
    ctx->id = 8968;
    ctx->coefficient = 2.8;
    ctx->flags = 230;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} 5C7rOxqM_ctx_t;

void init_ctx(5C7rOxqM_ctx_t *ctx) {
    ctx->id = 9069;
    ctx->coefficient = 7.2;
    ctx->flags = 2;
}

int compute_HnJxmXS(int input) {
    int result = input * 21;
    for(int i = 0; i < 12; i++) result += (i * 2);
    return result;
}

int compute_Gg6apQH(int input) {
    int result = input * 44;
    for(int i = 0; i < 7; i++) result += (i * 5);
    return result;
}

int compute_H8LRWjo(int input) {
    int result = input * 30;
    for(int i = 0; i < 9; i++) result += (i * 3);
    return result;
}

int compute_2UZarJc(int input) {
    int result = input * 13;
    for(int i = 0; i < 14; i++) result += (i * 1);
    return result;
}

int compute_K2Rjvxp(int input) {
    int result = input * 15;
    for(int i = 0; i < 13; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[77];
    double coefficient;
    uint8_t flags;
} IlKUbqHr_ctx_t;

void init_ctx(IlKUbqHr_ctx_t *ctx) {
    ctx->id = 4009;
    ctx->coefficient = 3.9;
    ctx->flags = 206;
}

int compute_srWwVKT(int input) {
    int result = input * 41;
    for(int i = 0; i < 14; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[67];
    double coefficient;
    uint8_t flags;
} gNdlGdnV_ctx_t;

void init_ctx(gNdlGdnV_ctx_t *ctx) {
    ctx->id = 6032;
    ctx->coefficient = 4.3;
    ctx->flags = 108;
}

int compute_CY1BBtW(int input) {
    int result = input * 67;
    for(int i = 0; i < 18; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[91];
    double coefficient;
    uint8_t flags;
} xiWNTwLa_ctx_t;

void init_ctx(xiWNTwLa_ctx_t *ctx) {
    ctx->id = 4806;
    ctx->coefficient = 7.6;
    ctx->flags = 26;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} BrDNv8Nb_ctx_t;

void init_ctx(BrDNv8Nb_ctx_t *ctx) {
    ctx->id = 9218;
    ctx->coefficient = 6.4;
    ctx->flags = 186;
}

int compute_MdS3KcB(int input) {
    int result = input * 6;
    for(int i = 0; i < 12; i++) result += (i * 5);
    return result;
}

int compute_UJvR7ZP(int input) {
    int result = input * 49;
    for(int i = 0; i < 6; i++) result += (i * 4);
    return result;
}

int compute_jTeyL3B(int input) {
    int result = input * 5;
    for(int i = 0; i < 10; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[66];
    double coefficient;
    uint8_t flags;
} iiwWEKuO_ctx_t;

void init_ctx(iiwWEKuO_ctx_t *ctx) {
    ctx->id = 6104;
    ctx->coefficient = 6.7;
    ctx->flags = 12;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} svwIALgF_ctx_t;

void init_ctx(svwIALgF_ctx_t *ctx) {
    ctx->id = 7544;
    ctx->coefficient = 5.6;
    ctx->flags = 22;
}

typedef struct {
    uint32_t id;
    char name[54];
    double coefficient;
    uint8_t flags;
} tWiWRz45_ctx_t;

void init_ctx(tWiWRz45_ctx_t *ctx) {
    ctx->id = 9867;
    ctx->coefficient = 4.9;
    ctx->flags = 24;
}

typedef struct {
    uint32_t id;
    char name[84];
    double coefficient;
    uint8_t flags;
} 1DolWtoF_ctx_t;

void init_ctx(1DolWtoF_ctx_t *ctx) {
    ctx->id = 9245;
    ctx->coefficient = 7.3;
    ctx->flags = 182;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} gb974NqN_ctx_t;

void init_ctx(gb974NqN_ctx_t *ctx) {
    ctx->id = 7656;
    ctx->coefficient = 7.7;
    ctx->flags = 140;
}

typedef struct {
    uint32_t id;
    char name[69];
    double coefficient;
    uint8_t flags;
} gE2qvVbC_ctx_t;

void init_ctx(gE2qvVbC_ctx_t *ctx) {
    ctx->id = 6871;
    ctx->coefficient = 4.4;
    ctx->flags = 205;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} eTOGmfwg_ctx_t;

void init_ctx(eTOGmfwg_ctx_t *ctx) {
    ctx->id = 2889;
    ctx->coefficient = 3.4;
    ctx->flags = 37;
}

int compute_H766EyY(int input) {
    int result = input * 34;
    for(int i = 0; i < 3; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[82];
    double coefficient;
    uint8_t flags;
} TK7DOX0T_ctx_t;

void init_ctx(TK7DOX0T_ctx_t *ctx) {
    ctx->id = 2106;
    ctx->coefficient = 5.8;
    ctx->flags = 252;
}

int compute_zacvxSl(int input) {
    int result = input * 67;
    for(int i = 0; i < 14; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[50];
    double coefficient;
    uint8_t flags;
} BTvjAS4p_ctx_t;

void init_ctx(BTvjAS4p_ctx_t *ctx) {
    ctx->id = 6808;
    ctx->coefficient = 8.6;
    ctx->flags = 61;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} XgZ5pHKK_ctx_t;

void init_ctx(XgZ5pHKK_ctx_t *ctx) {
    ctx->id = 3965;
    ctx->coefficient = 6.9;
    ctx->flags = 11;
}

int compute_w9Um3I3(int input) {
    int result = input * 3;
    for(int i = 0; i < 11; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[42];
    double coefficient;
    uint8_t flags;
} JbI9WfgK_ctx_t;

void init_ctx(JbI9WfgK_ctx_t *ctx) {
    ctx->id = 8083;
    ctx->coefficient = 9.0;
    ctx->flags = 17;
}

int compute_5a4ZwZr(int input) {
    int result = input * 22;
    for(int i = 0; i < 16; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[50];
    double coefficient;
    uint8_t flags;
} chhkL1x5_ctx_t;

void init_ctx(chhkL1x5_ctx_t *ctx) {
    ctx->id = 7761;
    ctx->coefficient = 2.0;
    ctx->flags = 117;
}

int compute_En8EFRt(int input) {
    int result = input * 27;
    for(int i = 0; i < 13; i++) result += (i * 5);
    return result;
}

int compute_sQQUs9W(int input) {
    int result = input * 56;
    for(int i = 0; i < 3; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[42];
    double coefficient;
    uint8_t flags;
} ePGUGfbe_ctx_t;

void init_ctx(ePGUGfbe_ctx_t *ctx) {
    ctx->id = 8079;
    ctx->coefficient = 3.9;
    ctx->flags = 10;
}

int compute_NreuSiW(int input) {
    int result = input * 31;
    for(int i = 0; i < 12; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[35];
    double coefficient;
    uint8_t flags;
} DL4gZw1Q_ctx_t;

void init_ctx(DL4gZw1Q_ctx_t *ctx) {
    ctx->id = 9202;
    ctx->coefficient = 1.2;
    ctx->flags = 164;
}

int compute_n493vaN(int input) {
    int result = input * 31;
    for(int i = 0; i < 5; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[57];
    double coefficient;
    uint8_t flags;
} TgMPGSR0_ctx_t;

void init_ctx(TgMPGSR0_ctx_t *ctx) {
    ctx->id = 9285;
    ctx->coefficient = 9.7;
    ctx->flags = 54;
}

typedef struct {
    uint32_t id;
    char name[62];
    double coefficient;
    uint8_t flags;
} pnvs4qEw_ctx_t;

void init_ctx(pnvs4qEw_ctx_t *ctx) {
    ctx->id = 5971;
    ctx->coefficient = 6.1;
    ctx->flags = 255;
}

int compute_ff8FdiN(int input) {
    int result = input * 52;
    for(int i = 0; i < 10; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[61];
    double coefficient;
    uint8_t flags;
} mMM4nRhs_ctx_t;

void init_ctx(mMM4nRhs_ctx_t *ctx) {
    ctx->id = 1269;
    ctx->coefficient = 3.4;
    ctx->flags = 199;
}

typedef struct {
    uint32_t id;
    char name[63];
    double coefficient;
    uint8_t flags;
} cbn4eYoE_ctx_t;

void init_ctx(cbn4eYoE_ctx_t *ctx) {
    ctx->id = 7454;
    ctx->coefficient = 3.2;
    ctx->flags = 192;
}

int compute_kL5g8Vm(int input) {
    int result = input * 68;
    for(int i = 0; i < 18; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[81];
    double coefficient;
    uint8_t flags;
} PCnbjFxy_ctx_t;

void init_ctx(PCnbjFxy_ctx_t *ctx) {
    ctx->id = 1555;
    ctx->coefficient = 4.2;
    ctx->flags = 243;
}

typedef struct {
    uint32_t id;
    char name[35];
    double coefficient;
    uint8_t flags;
} ZFi2t1i6_ctx_t;

void init_ctx(ZFi2t1i6_ctx_t *ctx) {
    ctx->id = 7238;
    ctx->coefficient = 4.9;
    ctx->flags = 237;
}

typedef struct {
    uint32_t id;
    char name[52];
    double coefficient;
    uint8_t flags;
} uOqh3cQW_ctx_t;

void init_ctx(uOqh3cQW_ctx_t *ctx) {
    ctx->id = 9497;
    ctx->coefficient = 0.9;
    ctx->flags = 59;
}

typedef struct {
    uint32_t id;
    char name[68];
    double coefficient;
    uint8_t flags;
} gx0JfvMX_ctx_t;

void init_ctx(gx0JfvMX_ctx_t *ctx) {
    ctx->id = 1097;
    ctx->coefficient = 7.7;
    ctx->flags = 145;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} nqM5epvR_ctx_t;

void init_ctx(nqM5epvR_ctx_t *ctx) {
    ctx->id = 4831;
    ctx->coefficient = 2.8;
    ctx->flags = 103;
}

typedef struct {
    uint32_t id;
    char name[96];
    double coefficient;
    uint8_t flags;
} iGofV8Zy_ctx_t;

void init_ctx(iGofV8Zy_ctx_t *ctx) {
    ctx->id = 3101;
    ctx->coefficient = 1.1;
    ctx->flags = 166;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} txQNuJBL_ctx_t;

void init_ctx(txQNuJBL_ctx_t *ctx) {
    ctx->id = 2468;
    ctx->coefficient = 7.1;
    ctx->flags = 253;
}

typedef struct {
    uint32_t id;
    char name[90];
    double coefficient;
    uint8_t flags;
} AFnnOjRV_ctx_t;

void init_ctx(AFnnOjRV_ctx_t *ctx) {
    ctx->id = 4379;
    ctx->coefficient = 5.1;
    ctx->flags = 124;
}

int compute_2thpTGg(int input) {
    int result = input * 27;
    for(int i = 0; i < 9; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} 3ElHsE5B_ctx_t;

void init_ctx(3ElHsE5B_ctx_t *ctx) {
    ctx->id = 3511;
    ctx->coefficient = 2.8;
    ctx->flags = 221;
}

int compute_PStwrZa(int input) {
    int result = input * 51;
    for(int i = 0; i < 4; i++) result += (i * 1);
    return result;
}

int compute_XifPt5H(int input) {
    int result = input * 53;
    for(int i = 0; i < 8; i++) result += (i * 5);
    return result;
}

int compute_IolvkN6(int input) {
    int result = input * 56;
    for(int i = 0; i < 4; i++) result += (i * 1);
    return result;
}

int compute_X8uO6Wk(int input) {
    int result = input * 47;
    for(int i = 0; i < 4; i++) result += (i * 2);
    return result;
}

int compute_KR614Fz(int input) {
    int result = input * 51;
    for(int i = 0; i < 14; i++) result += (i * 2);
    return result;
}

int compute_tcyBVg9(int input) {
    int result = input * 48;
    for(int i = 0; i < 7; i++) result += (i * 1);
    return result;
}

int compute_ZhV5ikA(int input) {
    int result = input * 27;
    for(int i = 0; i < 13; i++) result += (i * 5);
    return result;
}

int compute_vto9Ilb(int input) {
    int result = input * 24;
    for(int i = 0; i < 13; i++) result += (i * 2);
    return result;
}

int compute_yV7WO4P(int input) {
    int result = input * 60;
    for(int i = 0; i < 12; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} mGy8szOU_ctx_t;

void init_ctx(mGy8szOU_ctx_t *ctx) {
    ctx->id = 9216;
    ctx->coefficient = 6.6;
    ctx->flags = 42;
}

typedef struct {
    uint32_t id;
    char name[73];
    double coefficient;
    uint8_t flags;
} SIaz4TvS_ctx_t;

void init_ctx(SIaz4TvS_ctx_t *ctx) {
    ctx->id = 8500;
    ctx->coefficient = 2.2;
    ctx->flags = 131;
}

typedef struct {
    uint32_t id;
    char name[37];
    double coefficient;
    uint8_t flags;
} BifcPJc1_ctx_t;

void init_ctx(BifcPJc1_ctx_t *ctx) {
    ctx->id = 2335;
    ctx->coefficient = 6.0;
    ctx->flags = 112;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} n4Yqf1py_ctx_t;

void init_ctx(n4Yqf1py_ctx_t *ctx) {
    ctx->id = 5199;
    ctx->coefficient = 2.5;
    ctx->flags = 220;
}

int compute_TCo8Ixa(int input) {
    int result = input * 44;
    for(int i = 0; i < 3; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[81];
    double coefficient;
    uint8_t flags;
} tkFwBeyH_ctx_t;

void init_ctx(tkFwBeyH_ctx_t *ctx) {
    ctx->id = 9521;
    ctx->coefficient = 1.2;
    ctx->flags = 131;
}

typedef struct {
    uint32_t id;
    char name[92];
    double coefficient;
    uint8_t flags;
} Fu5Uhmkh_ctx_t;

void init_ctx(Fu5Uhmkh_ctx_t *ctx) {
    ctx->id = 2934;
    ctx->coefficient = 5.6;
    ctx->flags = 187;
}

typedef struct {
    uint32_t id;
    char name[82];
    double coefficient;
    uint8_t flags;
} jIFJJk2C_ctx_t;

void init_ctx(jIFJJk2C_ctx_t *ctx) {
    ctx->id = 7965;
    ctx->coefficient = 7.1;
    ctx->flags = 92;
}

int compute_2sejWOn(int input) {
    int result = input * 14;
    for(int i = 0; i < 16; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[92];
    double coefficient;
    uint8_t flags;
} r7rlXyOk_ctx_t;

void init_ctx(r7rlXyOk_ctx_t *ctx) {
    ctx->id = 4451;
    ctx->coefficient = 0.3;
    ctx->flags = 90;
}

typedef struct {
    uint32_t id;
    char name[74];
    double coefficient;
    uint8_t flags;
} QOy5fuw7_ctx_t;

void init_ctx(QOy5fuw7_ctx_t *ctx) {
    ctx->id = 4821;
    ctx->coefficient = 1.7;
    ctx->flags = 76;
}

int compute_e0fWM5c(int input) {
    int result = input * 42;
    for(int i = 0; i < 15; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[51];
    double coefficient;
    uint8_t flags;
} MBBn3fMA_ctx_t;

void init_ctx(MBBn3fMA_ctx_t *ctx) {
    ctx->id = 1453;
    ctx->coefficient = 8.8;
    ctx->flags = 138;
}

int compute_m2cDe6L(int input) {
    int result = input * 56;
    for(int i = 0; i < 4; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[64];
    double coefficient;
    uint8_t flags;
} 7e4GBZGe_ctx_t;

void init_ctx(7e4GBZGe_ctx_t *ctx) {
    ctx->id = 6192;
    ctx->coefficient = 3.1;
    ctx->flags = 158;
}

int compute_raIhaoO(int input) {
    int result = input * 20;
    for(int i = 0; i < 3; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} IYvM44C1_ctx_t;

void init_ctx(IYvM44C1_ctx_t *ctx) {
    ctx->id = 1580;
    ctx->coefficient = 8.8;
    ctx->flags = 21;
}

typedef struct {
    uint32_t id;
    char name[92];
    double coefficient;
    uint8_t flags;
} WW4TFkSY_ctx_t;

void init_ctx(WW4TFkSY_ctx_t *ctx) {
    ctx->id = 6996;
    ctx->coefficient = 0.8;
    ctx->flags = 149;
}

int compute_ucuGjTT(int input) {
    int result = input * 4;
    for(int i = 0; i < 12; i++) result += (i * 5);
    return result;
}

int compute_CkCHB6o(int input) {
    int result = input * 17;
    for(int i = 0; i < 3; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[81];
    double coefficient;
    uint8_t flags;
} zmYIjgOi_ctx_t;

void init_ctx(zmYIjgOi_ctx_t *ctx) {
    ctx->id = 7690;
    ctx->coefficient = 2.0;
    ctx->flags = 201;
}

int compute_IAsGS6h(int input) {
    int result = input * 42;
    for(int i = 0; i < 9; i++) result += (i * 1);
    return result;
}

int compute_e46G3jE(int input) {
    int result = input * 45;
    for(int i = 0; i < 18; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[96];
    double coefficient;
    uint8_t flags;
} XxH2WehW_ctx_t;

void init_ctx(XxH2WehW_ctx_t *ctx) {
    ctx->id = 9379;
    ctx->coefficient = 5.2;
    ctx->flags = 196;
}

int compute_NnDNtnZ(int input) {
    int result = input * 51;
    for(int i = 0; i < 4; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} sTgF148t_ctx_t;

void init_ctx(sTgF148t_ctx_t *ctx) {
    ctx->id = 3690;
    ctx->coefficient = 8.9;
    ctx->flags = 50;
}

int compute_0p13wPS(int input) {
    int result = input * 28;
    for(int i = 0; i < 11; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[56];
    double coefficient;
    uint8_t flags;
} 1yvDD8UQ_ctx_t;

void init_ctx(1yvDD8UQ_ctx_t *ctx) {
    ctx->id = 5084;
    ctx->coefficient = 7.3;
    ctx->flags = 158;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} CW6IwDvC_ctx_t;

void init_ctx(CW6IwDvC_ctx_t *ctx) {
    ctx->id = 9559;
    ctx->coefficient = 4.4;
    ctx->flags = 145;
}

typedef struct {
    uint32_t id;
    char name[72];
    double coefficient;
    uint8_t flags;
} kqKLTjPg_ctx_t;

void init_ctx(kqKLTjPg_ctx_t *ctx) {
    ctx->id = 1290;
    ctx->coefficient = 2.0;
    ctx->flags = 4;
}

int compute_QRK4EjS(int input) {
    int result = input * 41;
    for(int i = 0; i < 4; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[66];
    double coefficient;
    uint8_t flags;
} mgsYGvCk_ctx_t;

void init_ctx(mgsYGvCk_ctx_t *ctx) {
    ctx->id = 5800;
    ctx->coefficient = 8.0;
    ctx->flags = 7;
}

typedef struct {
    uint32_t id;
    char name[82];
    double coefficient;
    uint8_t flags;
} NeLM2WWX_ctx_t;

void init_ctx(NeLM2WWX_ctx_t *ctx) {
    ctx->id = 8311;
    ctx->coefficient = 6.6;
    ctx->flags = 108;
}

int compute_EC2PcSh(int input) {
    int result = input * 30;
    for(int i = 0; i < 17; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[50];
    double coefficient;
    uint8_t flags;
} szPgK1bX_ctx_t;

void init_ctx(szPgK1bX_ctx_t *ctx) {
    ctx->id = 1535;
    ctx->coefficient = 7.5;
    ctx->flags = 146;
}

int compute_obK33Dr(int input) {
    int result = input * 13;
    for(int i = 0; i < 3; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[72];
    double coefficient;
    uint8_t flags;
} SEpaLAYg_ctx_t;

void init_ctx(SEpaLAYg_ctx_t *ctx) {
    ctx->id = 1388;
    ctx->coefficient = 0.7;
    ctx->flags = 76;
}

typedef struct {
    uint32_t id;
    char name[84];
    double coefficient;
    uint8_t flags;
} sh1ka4Gm_ctx_t;

void init_ctx(sh1ka4Gm_ctx_t *ctx) {
    ctx->id = 4576;
    ctx->coefficient = 3.4;
    ctx->flags = 10;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} j0OHeSlR_ctx_t;

void init_ctx(j0OHeSlR_ctx_t *ctx) {
    ctx->id = 3133;
    ctx->coefficient = 7.6;
    ctx->flags = 41;
}

int compute_dDUQoEx(int input) {
    int result = input * 43;
    for(int i = 0; i < 13; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[39];
    double coefficient;
    uint8_t flags;
} eeFPLyFd_ctx_t;

void init_ctx(eeFPLyFd_ctx_t *ctx) {
    ctx->id = 6738;
    ctx->coefficient = 9.5;
    ctx->flags = 192;
}

typedef struct {
    uint32_t id;
    char name[65];
    double coefficient;
    uint8_t flags;
} KLlNdZYf_ctx_t;

void init_ctx(KLlNdZYf_ctx_t *ctx) {
    ctx->id = 6185;
    ctx->coefficient = 5.2;
    ctx->flags = 149;
}

typedef struct {
    uint32_t id;
    char name[80];
    double coefficient;
    uint8_t flags;
} fcy4H5Du_ctx_t;

void init_ctx(fcy4H5Du_ctx_t *ctx) {
    ctx->id = 7066;
    ctx->coefficient = 8.6;
    ctx->flags = 21;
}

int compute_vZeX2RT(int input) {
    int result = input * 23;
    for(int i = 0; i < 17; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[40];
    double coefficient;
    uint8_t flags;
} nKGQ4IVl_ctx_t;

void init_ctx(nKGQ4IVl_ctx_t *ctx) {
    ctx->id = 3663;
    ctx->coefficient = 5.3;
    ctx->flags = 64;
}

typedef struct {
    uint32_t id;
    char name[65];
    double coefficient;
    uint8_t flags;
} sKUWqWYX_ctx_t;

void init_ctx(sKUWqWYX_ctx_t *ctx) {
    ctx->id = 2686;
    ctx->coefficient = 6.4;
    ctx->flags = 34;
}

typedef struct {
    uint32_t id;
    char name[90];
    double coefficient;
    uint8_t flags;
} rAVe0lMS_ctx_t;

void init_ctx(rAVe0lMS_ctx_t *ctx) {
    ctx->id = 5819;
    ctx->coefficient = 9.1;
    ctx->flags = 192;
}

typedef struct {
    uint32_t id;
    char name[91];
    double coefficient;
    uint8_t flags;
} UFLOvRS5_ctx_t;

void init_ctx(UFLOvRS5_ctx_t *ctx) {
    ctx->id = 3830;
    ctx->coefficient = 3.3;
    ctx->flags = 117;
}

typedef struct {
    uint32_t id;
    char name[53];
    double coefficient;
    uint8_t flags;
} qE39XAT8_ctx_t;

void init_ctx(qE39XAT8_ctx_t *ctx) {
    ctx->id = 6317;
    ctx->coefficient = 1.5;
    ctx->flags = 81;
}

typedef struct {
    uint32_t id;
    char name[93];
    double coefficient;
    uint8_t flags;
} wNS7VLOD_ctx_t;

void init_ctx(wNS7VLOD_ctx_t *ctx) {
    ctx->id = 3074;
    ctx->coefficient = 5.6;
    ctx->flags = 100;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} 4maUQpUz_ctx_t;

void init_ctx(4maUQpUz_ctx_t *ctx) {
    ctx->id = 5915;
    ctx->coefficient = 0.4;
    ctx->flags = 208;
}

int compute_j1eIZwQ(int input) {
    int result = input * 56;
    for(int i = 0; i < 10; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[69];
    double coefficient;
    uint8_t flags;
} fg0hga0V_ctx_t;

void init_ctx(fg0hga0V_ctx_t *ctx) {
    ctx->id = 7334;
    ctx->coefficient = 7.8;
    ctx->flags = 139;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} WVVfljE3_ctx_t;

void init_ctx(WVVfljE3_ctx_t *ctx) {
    ctx->id = 6077;
    ctx->coefficient = 2.4;
    ctx->flags = 75;
}

typedef struct {
    uint32_t id;
    char name[77];
    double coefficient;
    uint8_t flags;
} PNy0154d_ctx_t;

void init_ctx(PNy0154d_ctx_t *ctx) {
    ctx->id = 1543;
    ctx->coefficient = 0.8;
    ctx->flags = 36;
}

typedef struct {
    uint32_t id;
    char name[55];
    double coefficient;
    uint8_t flags;
} caJckcfI_ctx_t;

void init_ctx(caJckcfI_ctx_t *ctx) {
    ctx->id = 6468;
    ctx->coefficient = 6.5;
    ctx->flags = 52;
}

typedef struct {
    uint32_t id;
    char name[90];
    double coefficient;
    uint8_t flags;
} loizff36_ctx_t;

void init_ctx(loizff36_ctx_t *ctx) {
    ctx->id = 5049;
    ctx->coefficient = 4.9;
    ctx->flags = 60;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} KvXb03XK_ctx_t;

void init_ctx(KvXb03XK_ctx_t *ctx) {
    ctx->id = 6013;
    ctx->coefficient = 5.7;
    ctx->flags = 122;
}

typedef struct {
    uint32_t id;
    char name[38];
    double coefficient;
    uint8_t flags;
} u6nbbIwd_ctx_t;

void init_ctx(u6nbbIwd_ctx_t *ctx) {
    ctx->id = 7937;
    ctx->coefficient = 2.9;
    ctx->flags = 179;
}

typedef struct {
    uint32_t id;
    char name[68];
    double coefficient;
    uint8_t flags;
} zmk6JgOW_ctx_t;

void init_ctx(zmk6JgOW_ctx_t *ctx) {
    ctx->id = 9237;
    ctx->coefficient = 6.9;
    ctx->flags = 249;
}

typedef struct {
    uint32_t id;
    char name[72];
    double coefficient;
    uint8_t flags;
} a2IyTHn3_ctx_t;

void init_ctx(a2IyTHn3_ctx_t *ctx) {
    ctx->id = 8862;
    ctx->coefficient = 4.8;
    ctx->flags = 146;
}

int compute_9cvUNwc(int input) {
    int result = input * 34;
    for(int i = 0; i < 16; i++) result += (i * 2);
    return result;
}

int compute_YawZFLf(int input) {
    int result = input * 66;
    for(int i = 0; i < 18; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} yIyr2JTZ_ctx_t;

void init_ctx(yIyr2JTZ_ctx_t *ctx) {
    ctx->id = 3664;
    ctx->coefficient = 8.4;
    ctx->flags = 8;
}

typedef struct {
    uint32_t id;
    char name[37];
    double coefficient;
    uint8_t flags;
} 43EowW6t_ctx_t;

void init_ctx(43EowW6t_ctx_t *ctx) {
    ctx->id = 8789;
    ctx->coefficient = 7.3;
    ctx->flags = 128;
}

typedef struct {
    uint32_t id;
    char name[58];
    double coefficient;
    uint8_t flags;
} QSxsfwDw_ctx_t;

void init_ctx(QSxsfwDw_ctx_t *ctx) {
    ctx->id = 4803;
    ctx->coefficient = 1.5;
    ctx->flags = 58;
}

typedef struct {
    uint32_t id;
    char name[83];
    double coefficient;
    uint8_t flags;
} q0RyBkRT_ctx_t;

void init_ctx(q0RyBkRT_ctx_t *ctx) {
    ctx->id = 2168;
    ctx->coefficient = 4.3;
    ctx->flags = 56;
}

typedef struct {
    uint32_t id;
    char name[62];
    double coefficient;
    uint8_t flags;
} jFhxWkay_ctx_t;

void init_ctx(jFhxWkay_ctx_t *ctx) {
    ctx->id = 7442;
    ctx->coefficient = 6.7;
    ctx->flags = 122;
}

int compute_vJUqnYH(int input) {
    int result = input * 46;
    for(int i = 0; i < 11; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[62];
    double coefficient;
    uint8_t flags;
} w9i6xRUf_ctx_t;

void init_ctx(w9i6xRUf_ctx_t *ctx) {
    ctx->id = 8147;
    ctx->coefficient = 0.7;
    ctx->flags = 75;
}

typedef struct {
    uint32_t id;
    char name[76];
    double coefficient;
    uint8_t flags;
} gMprwwTt_ctx_t;

void init_ctx(gMprwwTt_ctx_t *ctx) {
    ctx->id = 1758;
    ctx->coefficient = 9.5;
    ctx->flags = 69;
}

int compute_t1ReXLQ(int input) {
    int result = input * 4;
    for(int i = 0; i < 8; i++) result += (i * 5);
    return result;
}

int compute_DQjyEJL(int input) {
    int result = input * 58;
    for(int i = 0; i < 18; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[80];
    double coefficient;
    uint8_t flags;
} aEc76NYu_ctx_t;

void init_ctx(aEc76NYu_ctx_t *ctx) {
    ctx->id = 2085;
    ctx->coefficient = 5.4;
    ctx->flags = 59;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} VXsQY69m_ctx_t;

void init_ctx(VXsQY69m_ctx_t *ctx) {
    ctx->id = 7153;
    ctx->coefficient = 8.9;
    ctx->flags = 127;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} OTPqtiun_ctx_t;

void init_ctx(OTPqtiun_ctx_t *ctx) {
    ctx->id = 6673;
    ctx->coefficient = 5.1;
    ctx->flags = 3;
}

typedef struct {
    uint32_t id;
    char name[45];
    double coefficient;
    uint8_t flags;
} aAd119Wd_ctx_t;

void init_ctx(aAd119Wd_ctx_t *ctx) {
    ctx->id = 9588;
    ctx->coefficient = 4.1;
    ctx->flags = 73;
}

int compute_cU1sJBR(int input) {
    int result = input * 42;
    for(int i = 0; i < 14; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[72];
    double coefficient;
    uint8_t flags;
} mCVLFWf7_ctx_t;

void init_ctx(mCVLFWf7_ctx_t *ctx) {
    ctx->id = 2069;
    ctx->coefficient = 4.3;
    ctx->flags = 44;
}

typedef struct {
    uint32_t id;
    char name[66];
    double coefficient;
    uint8_t flags;
} g6X0exjN_ctx_t;

void init_ctx(g6X0exjN_ctx_t *ctx) {
    ctx->id = 7176;
    ctx->coefficient = 2.9;
    ctx->flags = 41;
}

int compute_vxIOg03(int input) {
    int result = input * 41;
    for(int i = 0; i < 6; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[37];
    double coefficient;
    uint8_t flags;
} N8oGUMey_ctx_t;

void init_ctx(N8oGUMey_ctx_t *ctx) {
    ctx->id = 1191;
    ctx->coefficient = 1.8;
    ctx->flags = 45;
}

typedef struct {
    uint32_t id;
    char name[60];
    double coefficient;
    uint8_t flags;
} qJwtYySC_ctx_t;

void init_ctx(qJwtYySC_ctx_t *ctx) {
    ctx->id = 5111;
    ctx->coefficient = 6.1;
    ctx->flags = 73;
}

typedef struct {
    uint32_t id;
    char name[52];
    double coefficient;
    uint8_t flags;
} JZsdruRc_ctx_t;

void init_ctx(JZsdruRc_ctx_t *ctx) {
    ctx->id = 6278;
    ctx->coefficient = 1.3;
    ctx->flags = 93;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} DlQnAax1_ctx_t;

void init_ctx(DlQnAax1_ctx_t *ctx) {
    ctx->id = 2526;
    ctx->coefficient = 0.3;
    ctx->flags = 65;
}

int compute_qskjrXZ(int input) {
    int result = input * 38;
    for(int i = 0; i < 17; i++) result += (i * 2);
    return result;
}

int compute_5WxtojF(int input) {
    int result = input * 43;
    for(int i = 0; i < 9; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[59];
    double coefficient;
    uint8_t flags;
} aowVjDV8_ctx_t;

void init_ctx(aowVjDV8_ctx_t *ctx) {
    ctx->id = 3107;
    ctx->coefficient = 3.2;
    ctx->flags = 98;
}

typedef struct {
    uint32_t id;
    char name[74];
    double coefficient;
    uint8_t flags;
} 5cOiPqQ9_ctx_t;

void init_ctx(5cOiPqQ9_ctx_t *ctx) {
    ctx->id = 1546;
    ctx->coefficient = 2.1;
    ctx->flags = 31;
}

typedef struct {
    uint32_t id;
    char name[53];
    double coefficient;
    uint8_t flags;
} K5Okj5iv_ctx_t;

void init_ctx(K5Okj5iv_ctx_t *ctx) {
    ctx->id = 9153;
    ctx->coefficient = 4.1;
    ctx->flags = 149;
}

int compute_6dZIXTP(int input) {
    int result = input * 64;
    for(int i = 0; i < 13; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} LkL17sjB_ctx_t;

void init_ctx(LkL17sjB_ctx_t *ctx) {
    ctx->id = 2607;
    ctx->coefficient = 3.7;
    ctx->flags = 1;
}

typedef struct {
    uint32_t id;
    char name[59];
    double coefficient;
    uint8_t flags;
} VSFhHSh3_ctx_t;

void init_ctx(VSFhHSh3_ctx_t *ctx) {
    ctx->id = 6434;
    ctx->coefficient = 5.4;
    ctx->flags = 19;
}

int compute_gpXVP7N(int input) {
    int result = input * 26;
    for(int i = 0; i < 8; i++) result += (i * 2);
    return result;
}

int compute_uYKtp5o(int input) {
    int result = input * 12;
    for(int i = 0; i < 13; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[49];
    double coefficient;
    uint8_t flags;
} 26mBmo3k_ctx_t;

void init_ctx(26mBmo3k_ctx_t *ctx) {
    ctx->id = 2503;
    ctx->coefficient = 1.7;
    ctx->flags = 210;
}

int compute_JSLB6aJ(int input) {
    int result = input * 45;
    for(int i = 0; i < 10; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[83];
    double coefficient;
    uint8_t flags;
} SKz2njaG_ctx_t;

void init_ctx(SKz2njaG_ctx_t *ctx) {
    ctx->id = 4601;
    ctx->coefficient = 7.2;
    ctx->flags = 207;
}

typedef struct {
    uint32_t id;
    char name[62];
    double coefficient;
    uint8_t flags;
} TURyHJ6M_ctx_t;

void init_ctx(TURyHJ6M_ctx_t *ctx) {
    ctx->id = 3757;
    ctx->coefficient = 2.8;
    ctx->flags = 242;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} axX5WsEU_ctx_t;

void init_ctx(axX5WsEU_ctx_t *ctx) {
    ctx->id = 3205;
    ctx->coefficient = 1.6;
    ctx->flags = 129;
}

typedef struct {
    uint32_t id;
    char name[52];
    double coefficient;
    uint8_t flags;
} GJ79i5R3_ctx_t;

void init_ctx(GJ79i5R3_ctx_t *ctx) {
    ctx->id = 8427;
    ctx->coefficient = 4.2;
    ctx->flags = 218;
}

int compute_aJpdu6R(int input) {
    int result = input * 66;
    for(int i = 0; i < 11; i++) result += (i * 1);
    return result;
}

int compute_rU6q2QJ(int input) {
    int result = input * 22;
    for(int i = 0; i < 6; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[74];
    double coefficient;
    uint8_t flags;
} IcGGx2ZJ_ctx_t;

void init_ctx(IcGGx2ZJ_ctx_t *ctx) {
    ctx->id = 8295;
    ctx->coefficient = 7.0;
    ctx->flags = 174;
}

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} yDctdp0m_ctx_t;

void init_ctx(yDctdp0m_ctx_t *ctx) {
    ctx->id = 2219;
    ctx->coefficient = 7.2;
    ctx->flags = 48;
}

typedef struct {
    uint32_t id;
    char name[68];
    double coefficient;
    uint8_t flags;
} v3HIqjQB_ctx_t;

void init_ctx(v3HIqjQB_ctx_t *ctx) {
    ctx->id = 2257;
    ctx->coefficient = 7.0;
    ctx->flags = 91;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} ZHxGzaQ0_ctx_t;

void init_ctx(ZHxGzaQ0_ctx_t *ctx) {
    ctx->id = 7746;
    ctx->coefficient = 3.2;
    ctx->flags = 252;
}

typedef struct {
    uint32_t id;
    char name[61];
    double coefficient;
    uint8_t flags;
} iEq2SrYk_ctx_t;

void init_ctx(iEq2SrYk_ctx_t *ctx) {
    ctx->id = 7095;
    ctx->coefficient = 3.5;
    ctx->flags = 125;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} fb60jVQ5_ctx_t;

void init_ctx(fb60jVQ5_ctx_t *ctx) {
    ctx->id = 3083;
    ctx->coefficient = 5.2;
    ctx->flags = 234;
}

typedef struct {
    uint32_t id;
    char name[82];
    double coefficient;
    uint8_t flags;
} 45WrVeeR_ctx_t;

void init_ctx(45WrVeeR_ctx_t *ctx) {
    ctx->id = 6946;
    ctx->coefficient = 0.4;
    ctx->flags = 57;
}

typedef struct {
    uint32_t id;
    char name[45];
    double coefficient;
    uint8_t flags;
} jLkra8kX_ctx_t;

void init_ctx(jLkra8kX_ctx_t *ctx) {
    ctx->id = 3925;
    ctx->coefficient = 5.4;
    ctx->flags = 189;
}

typedef struct {
    uint32_t id;
    char name[91];
    double coefficient;
    uint8_t flags;
} wFAJXsXK_ctx_t;

void init_ctx(wFAJXsXK_ctx_t *ctx) {
    ctx->id = 4061;
    ctx->coefficient = 4.2;
    ctx->flags = 159;
}

int compute_fGyM1BG(int input) {
    int result = input * 60;
    for(int i = 0; i < 4; i++) result += (i * 4);
    return result;
}

int compute_ohnzfsl(int input) {
    int result = input * 65;
    for(int i = 0; i < 16; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} uop0AxrP_ctx_t;

void init_ctx(uop0AxrP_ctx_t *ctx) {
    ctx->id = 4219;
    ctx->coefficient = 6.7;
    ctx->flags = 105;
}

int compute_csGUJCL(int input) {
    int result = input * 8;
    for(int i = 0; i < 8; i++) result += (i * 2);
    return result;
}

int compute_xpA4wBM(int input) {
    int result = input * 35;
    for(int i = 0; i < 17; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} S9JioPg3_ctx_t;

void init_ctx(S9JioPg3_ctx_t *ctx) {
    ctx->id = 3414;
    ctx->coefficient = 0.3;
    ctx->flags = 30;
}

int compute_m9lhAiC(int input) {
    int result = input * 14;
    for(int i = 0; i < 14; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[67];
    double coefficient;
    uint8_t flags;
} pmOMz11c_ctx_t;

void init_ctx(pmOMz11c_ctx_t *ctx) {
    ctx->id = 3866;
    ctx->coefficient = 7.8;
    ctx->flags = 46;
}

int compute_AUerN5p(int input) {
    int result = input * 58;
    for(int i = 0; i < 17; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} TZGwT26w_ctx_t;

void init_ctx(TZGwT26w_ctx_t *ctx) {
    ctx->id = 7078;
    ctx->coefficient = 9.4;
    ctx->flags = 211;
}

int compute_mzpSHHU(int input) {
    int result = input * 21;
    for(int i = 0; i < 13; i++) result += (i * 3);
    return result;
}

int compute_xwwFO1q(int input) {
    int result = input * 62;
    for(int i = 0; i < 4; i++) result += (i * 3);
    return result;
}

int compute_4EKDEiX(int input) {
    int result = input * 7;
    for(int i = 0; i < 16; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[53];
    double coefficient;
    uint8_t flags;
} 9g1LN6kl_ctx_t;

void init_ctx(9g1LN6kl_ctx_t *ctx) {
    ctx->id = 5346;
    ctx->coefficient = 7.7;
    ctx->flags = 170;
}

typedef struct {
    uint32_t id;
    char name[46];
    double coefficient;
    uint8_t flags;
} mImaIeab_ctx_t;

void init_ctx(mImaIeab_ctx_t *ctx) {
    ctx->id = 7227;
    ctx->coefficient = 0.4;
    ctx->flags = 190;
}

int compute_SbyPifi(int input) {
    int result = input * 42;
    for(int i = 0; i < 5; i++) result += (i * 3);
    return result;
}

int compute_jc0EhJI(int input) {
    int result = input * 37;
    for(int i = 0; i < 11; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} oR6tFMaN_ctx_t;

void init_ctx(oR6tFMaN_ctx_t *ctx) {
    ctx->id = 6774;
    ctx->coefficient = 4.7;
    ctx->flags = 40;
}

typedef struct {
    uint32_t id;
    char name[52];
    double coefficient;
    uint8_t flags;
} u6cBCll6_ctx_t;

void init_ctx(u6cBCll6_ctx_t *ctx) {
    ctx->id = 3676;
    ctx->coefficient = 2.9;
    ctx->flags = 163;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} Lv8FTpaB_ctx_t;

void init_ctx(Lv8FTpaB_ctx_t *ctx) {
    ctx->id = 4164;
    ctx->coefficient = 9.7;
    ctx->flags = 0;
}

typedef struct {
    uint32_t id;
    char name[41];
    double coefficient;
    uint8_t flags;
} RrxFiENO_ctx_t;

void init_ctx(RrxFiENO_ctx_t *ctx) {
    ctx->id = 9686;
    ctx->coefficient = 8.7;
    ctx->flags = 144;
}

typedef struct {
    uint32_t id;
    char name[53];
    double coefficient;
    uint8_t flags;
} 1k8o5GJU_ctx_t;

void init_ctx(1k8o5GJU_ctx_t *ctx) {
    ctx->id = 3610;
    ctx->coefficient = 9.5;
    ctx->flags = 213;
}

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} B3uBz6uH_ctx_t;

void init_ctx(B3uBz6uH_ctx_t *ctx) {
    ctx->id = 3821;
    ctx->coefficient = 9.9;
    ctx->flags = 96;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} B26H2ZUz_ctx_t;

void init_ctx(B26H2ZUz_ctx_t *ctx) {
    ctx->id = 7058;
    ctx->coefficient = 0.4;
    ctx->flags = 224;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} WeiM7Ovm_ctx_t;

void init_ctx(WeiM7Ovm_ctx_t *ctx) {
    ctx->id = 3159;
    ctx->coefficient = 0.5;
    ctx->flags = 47;
}

typedef struct {
    uint32_t id;
    char name[91];
    double coefficient;
    uint8_t flags;
} SE58xyEw_ctx_t;

void init_ctx(SE58xyEw_ctx_t *ctx) {
    ctx->id = 6660;
    ctx->coefficient = 10.0;
    ctx->flags = 166;
}

typedef struct {
    uint32_t id;
    char name[51];
    double coefficient;
    uint8_t flags;
} BQWhdR1b_ctx_t;

void init_ctx(BQWhdR1b_ctx_t *ctx) {
    ctx->id = 9732;
    ctx->coefficient = 10.0;
    ctx->flags = 196;
}

typedef struct {
    uint32_t id;
    char name[91];
    double coefficient;
    uint8_t flags;
} 1NQjVn0Z_ctx_t;

void init_ctx(1NQjVn0Z_ctx_t *ctx) {
    ctx->id = 1733;
    ctx->coefficient = 0.3;
    ctx->flags = 17;
}

typedef struct {
    uint32_t id;
    char name[41];
    double coefficient;
    uint8_t flags;
} 7Cl7KOKo_ctx_t;

void init_ctx(7Cl7KOKo_ctx_t *ctx) {
    ctx->id = 2336;
    ctx->coefficient = 3.3;
    ctx->flags = 221;
}

typedef struct {
    uint32_t id;
    char name[59];
    double coefficient;
    uint8_t flags;
} hIGTa3Be_ctx_t;

void init_ctx(hIGTa3Be_ctx_t *ctx) {
    ctx->id = 8010;
    ctx->coefficient = 4.5;
    ctx->flags = 246;
}

int compute_tL26ido(int input) {
    int result = input * 69;
    for(int i = 0; i < 13; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[61];
    double coefficient;
    uint8_t flags;
} XjoIjjzc_ctx_t;

void init_ctx(XjoIjjzc_ctx_t *ctx) {
    ctx->id = 4592;
    ctx->coefficient = 7.1;
    ctx->flags = 155;
}

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} DpTPMT0s_ctx_t;

void init_ctx(DpTPMT0s_ctx_t *ctx) {
    ctx->id = 9771;
    ctx->coefficient = 4.3;
    ctx->flags = 107;
}

typedef struct {
    uint32_t id;
    char name[95];
    double coefficient;
    uint8_t flags;
} ichxSTkI_ctx_t;

void init_ctx(ichxSTkI_ctx_t *ctx) {
    ctx->id = 7567;
    ctx->coefficient = 7.9;
    ctx->flags = 22;
}

int compute_1V1n0kn(int input) {
    int result = input * 51;
    for(int i = 0; i < 17; i++) result += (i * 4);
    return result;
}

int compute_i1JAKCP(int input) {
    int result = input * 34;
    for(int i = 0; i < 3; i++) result += (i * 4);
    return result;
}

int compute_W5uygMU(int input) {
    int result = input * 54;
    for(int i = 0; i < 3; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[84];
    double coefficient;
    uint8_t flags;
} uT8fSXdj_ctx_t;

void init_ctx(uT8fSXdj_ctx_t *ctx) {
    ctx->id = 5815;
    ctx->coefficient = 5.7;
    ctx->flags = 66;
}

typedef struct {
    uint32_t id;
    char name[46];
    double coefficient;
    uint8_t flags;
} IuhIatKp_ctx_t;

void init_ctx(IuhIatKp_ctx_t *ctx) {
    ctx->id = 1651;
    ctx->coefficient = 3.2;
    ctx->flags = 130;
}

int compute_oVbAtfI(int input) {
    int result = input * 45;
    for(int i = 0; i < 15; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[47];
    double coefficient;
    uint8_t flags;
} vElZ7ivb_ctx_t;

void init_ctx(vElZ7ivb_ctx_t *ctx) {
    ctx->id = 8355;
    ctx->coefficient = 4.8;
    ctx->flags = 222;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} A0s2V5iR_ctx_t;

void init_ctx(A0s2V5iR_ctx_t *ctx) {
    ctx->id = 2370;
    ctx->coefficient = 10.0;
    ctx->flags = 23;
}

typedef struct {
    uint32_t id;
    char name[50];
    double coefficient;
    uint8_t flags;
} UJKbPOCx_ctx_t;

void init_ctx(UJKbPOCx_ctx_t *ctx) {
    ctx->id = 2160;
    ctx->coefficient = 8.3;
    ctx->flags = 119;
}

typedef struct {
    uint32_t id;
    char name[42];
    double coefficient;
    uint8_t flags;
} AzEUqQRa_ctx_t;

void init_ctx(AzEUqQRa_ctx_t *ctx) {
    ctx->id = 5452;
    ctx->coefficient = 7.8;
    ctx->flags = 215;
}

typedef struct {
    uint32_t id;
    char name[83];
    double coefficient;
    uint8_t flags;
} rOwjwGe6_ctx_t;

void init_ctx(rOwjwGe6_ctx_t *ctx) {
    ctx->id = 4699;
    ctx->coefficient = 6.0;
    ctx->flags = 233;
}

typedef struct {
    uint32_t id;
    char name[90];
    double coefficient;
    uint8_t flags;
} csHLtX47_ctx_t;

void init_ctx(csHLtX47_ctx_t *ctx) {
    ctx->id = 9629;
    ctx->coefficient = 9.6;
    ctx->flags = 226;
}

int compute_517KT0B(int input) {
    int result = input * 70;
    for(int i = 0; i < 13; i++) result += (i * 3);
    return result;
}

int compute_1sJ7NVr(int input) {
    int result = input * 39;
    for(int i = 0; i < 8; i++) result += (i * 5);
    return result;
}

int compute_XFGXEIR(int input) {
    int result = input * 56;
    for(int i = 0; i < 4; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} DFfF4P2U_ctx_t;

void init_ctx(DFfF4P2U_ctx_t *ctx) {
    ctx->id = 8725;
    ctx->coefficient = 4.5;
    ctx->flags = 99;
}

typedef struct {
    uint32_t id;
    char name[53];
    double coefficient;
    uint8_t flags;
} B4L5Tgmo_ctx_t;

void init_ctx(B4L5Tgmo_ctx_t *ctx) {
    ctx->id = 6973;
    ctx->coefficient = 4.4;
    ctx->flags = 167;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} xCQsPc6K_ctx_t;

void init_ctx(xCQsPc6K_ctx_t *ctx) {
    ctx->id = 6922;
    ctx->coefficient = 0.6;
    ctx->flags = 86;
}

int compute_hXLUlxz(int input) {
    int result = input * 17;
    for(int i = 0; i < 14; i++) result += (i * 2);
    return result;
}

int compute_4YeYoEe(int input) {
    int result = input * 40;
    for(int i = 0; i < 9; i++) result += (i * 1);
    return result;
}

int compute_jRXK3bJ(int input) {
    int result = input * 60;
    for(int i = 0; i < 12; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[46];
    double coefficient;
    uint8_t flags;
} 4MWObPcy_ctx_t;

void init_ctx(4MWObPcy_ctx_t *ctx) {
    ctx->id = 6187;
    ctx->coefficient = 8.7;
    ctx->flags = 42;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} oUcllM9J_ctx_t;

void init_ctx(oUcllM9J_ctx_t *ctx) {
    ctx->id = 4788;
    ctx->coefficient = 5.1;
    ctx->flags = 236;
}

int compute_ezXQDs9(int input) {
    int result = input * 40;
    for(int i = 0; i < 12; i++) result += (i * 5);
    return result;
}

int compute_PzCaLtx(int input) {
    int result = input * 47;
    for(int i = 0; i < 3; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[79];
    double coefficient;
    uint8_t flags;
} pPWyQ9W3_ctx_t;

void init_ctx(pPWyQ9W3_ctx_t *ctx) {
    ctx->id = 9623;
    ctx->coefficient = 9.3;
    ctx->flags = 193;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} ajgo2qSR_ctx_t;

void init_ctx(ajgo2qSR_ctx_t *ctx) {
    ctx->id = 1252;
    ctx->coefficient = 7.2;
    ctx->flags = 207;
}

typedef struct {
    uint32_t id;
    char name[73];
    double coefficient;
    uint8_t flags;
} 39f3WCMp_ctx_t;

void init_ctx(39f3WCMp_ctx_t *ctx) {
    ctx->id = 2665;
    ctx->coefficient = 6.3;
    ctx->flags = 69;
}

typedef struct {
    uint32_t id;
    char name[66];
    double coefficient;
    uint8_t flags;
} Fi6A2ei7_ctx_t;

void init_ctx(Fi6A2ei7_ctx_t *ctx) {
    ctx->id = 2284;
    ctx->coefficient = 7.3;
    ctx->flags = 148;
}

typedef struct {
    uint32_t id;
    char name[40];
    double coefficient;
    uint8_t flags;
} dGgzI8Jq_ctx_t;

void init_ctx(dGgzI8Jq_ctx_t *ctx) {
    ctx->id = 2357;
    ctx->coefficient = 3.2;
    ctx->flags = 178;
}

typedef struct {
    uint32_t id;
    char name[55];
    double coefficient;
    uint8_t flags;
} VbrSgFh2_ctx_t;

void init_ctx(VbrSgFh2_ctx_t *ctx) {
    ctx->id = 3399;
    ctx->coefficient = 8.5;
    ctx->flags = 131;
}

int compute_iIyBhkw(int input) {
    int result = input * 14;
    for(int i = 0; i < 17; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[42];
    double coefficient;
    uint8_t flags;
} LWPk7P1Q_ctx_t;

void init_ctx(LWPk7P1Q_ctx_t *ctx) {
    ctx->id = 3640;
    ctx->coefficient = 2.2;
    ctx->flags = 99;
}

typedef struct {
    uint32_t id;
    char name[45];
    double coefficient;
    uint8_t flags;
} 3OubInQ4_ctx_t;

void init_ctx(3OubInQ4_ctx_t *ctx) {
    ctx->id = 7290;
    ctx->coefficient = 9.7;
    ctx->flags = 58;
}

typedef struct {
    uint32_t id;
    char name[50];
    double coefficient;
    uint8_t flags;
} HTil9DR1_ctx_t;

void init_ctx(HTil9DR1_ctx_t *ctx) {
    ctx->id = 3169;
    ctx->coefficient = 7.3;
    ctx->flags = 200;
}

int compute_hIInnM1(int input) {
    int result = input * 20;
    for(int i = 0; i < 14; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[83];
    double coefficient;
    uint8_t flags;
} bBmIKQEx_ctx_t;

void init_ctx(bBmIKQEx_ctx_t *ctx) {
    ctx->id = 6585;
    ctx->coefficient = 2.1;
    ctx->flags = 206;
}

typedef struct {
    uint32_t id;
    char name[38];
    double coefficient;
    uint8_t flags;
} 8VRRgTgI_ctx_t;

void init_ctx(8VRRgTgI_ctx_t *ctx) {
    ctx->id = 3340;
    ctx->coefficient = 1.7;
    ctx->flags = 221;
}

int compute_qCjVXDC(int input) {
    int result = input * 35;
    for(int i = 0; i < 16; i++) result += (i * 5);
    return result;
}

int compute_vywiBSm(int input) {
    int result = input * 60;
    for(int i = 0; i < 14; i++) result += (i * 2);
    return result;
}

int compute_zIIgGjl(int input) {
    int result = input * 23;
    for(int i = 0; i < 14; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[46];
    double coefficient;
    uint8_t flags;
} J07Ga9Qf_ctx_t;

void init_ctx(J07Ga9Qf_ctx_t *ctx) {
    ctx->id = 4333;
    ctx->coefficient = 4.5;
    ctx->flags = 134;
}

typedef struct {
    uint32_t id;
    char name[93];
    double coefficient;
    uint8_t flags;
} fBjb6wjQ_ctx_t;

void init_ctx(fBjb6wjQ_ctx_t *ctx) {
    ctx->id = 4349;
    ctx->coefficient = 4.7;
    ctx->flags = 236;
}

typedef struct {
    uint32_t id;
    char name[64];
    double coefficient;
    uint8_t flags;
} m5wAAyiP_ctx_t;

void init_ctx(m5wAAyiP_ctx_t *ctx) {
    ctx->id = 5515;
    ctx->coefficient = 4.2;
    ctx->flags = 246;
}

typedef struct {
    uint32_t id;
    char name[51];
    double coefficient;
    uint8_t flags;
} gFUgqCzE_ctx_t;

void init_ctx(gFUgqCzE_ctx_t *ctx) {
    ctx->id = 3735;
    ctx->coefficient = 4.4;
    ctx->flags = 193;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} dXifclwa_ctx_t;

void init_ctx(dXifclwa_ctx_t *ctx) {
    ctx->id = 8752;
    ctx->coefficient = 4.1;
    ctx->flags = 167;
}

typedef struct {
    uint32_t id;
    char name[75];
    double coefficient;
    uint8_t flags;
} ytsmhyRt_ctx_t;

void init_ctx(ytsmhyRt_ctx_t *ctx) {
    ctx->id = 1591;
    ctx->coefficient = 0.8;
    ctx->flags = 96;
}

int compute_YfVkdm0(int input) {
    int result = input * 59;
    for(int i = 0; i < 12; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} Lc8Lqn1B_ctx_t;

void init_ctx(Lc8Lqn1B_ctx_t *ctx) {
    ctx->id = 7096;
    ctx->coefficient = 5.3;
    ctx->flags = 35;
}

typedef struct {
    uint32_t id;
    char name[46];
    double coefficient;
    uint8_t flags;
} L7QZ5vNe_ctx_t;

void init_ctx(L7QZ5vNe_ctx_t *ctx) {
    ctx->id = 3851;
    ctx->coefficient = 3.4;
    ctx->flags = 149;
}

typedef struct {
    uint32_t id;
    char name[79];
    double coefficient;
    uint8_t flags;
} AS2C7Qev_ctx_t;

void init_ctx(AS2C7Qev_ctx_t *ctx) {
    ctx->id = 8907;
    ctx->coefficient = 9.8;
    ctx->flags = 131;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} sdVBkQaG_ctx_t;

void init_ctx(sdVBkQaG_ctx_t *ctx) {
    ctx->id = 6889;
    ctx->coefficient = 1.5;
    ctx->flags = 68;
}

int compute_AhmxpIX(int input) {
    int result = input * 29;
    for(int i = 0; i < 8; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[33];
    double coefficient;
    uint8_t flags;
} OTgIHV1i_ctx_t;

void init_ctx(OTgIHV1i_ctx_t *ctx) {
    ctx->id = 1210;
    ctx->coefficient = 4.7;
    ctx->flags = 10;
}

typedef struct {
    uint32_t id;
    char name[43];
    double coefficient;
    uint8_t flags;
} lZos79js_ctx_t;

void init_ctx(lZos79js_ctx_t *ctx) {
    ctx->id = 1874;
    ctx->coefficient = 9.2;
    ctx->flags = 114;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} uFxktN2O_ctx_t;

void init_ctx(uFxktN2O_ctx_t *ctx) {
    ctx->id = 1632;
    ctx->coefficient = 7.2;
    ctx->flags = 203;
}

int compute_HLQfNal(int input) {
    int result = input * 23;
    for(int i = 0; i < 9; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[35];
    double coefficient;
    uint8_t flags;
} qhn1SvOQ_ctx_t;

void init_ctx(qhn1SvOQ_ctx_t *ctx) {
    ctx->id = 2011;
    ctx->coefficient = 8.1;
    ctx->flags = 120;
}

typedef struct {
    uint32_t id;
    char name[90];
    double coefficient;
    uint8_t flags;
} GbyeUi7G_ctx_t;

void init_ctx(GbyeUi7G_ctx_t *ctx) {
    ctx->id = 9005;
    ctx->coefficient = 3.0;
    ctx->flags = 173;
}

typedef struct {
    uint32_t id;
    char name[61];
    double coefficient;
    uint8_t flags;
} YQVfMV4Y_ctx_t;

void init_ctx(YQVfMV4Y_ctx_t *ctx) {
    ctx->id = 7011;
    ctx->coefficient = 5.7;
    ctx->flags = 32;
}

typedef struct {
    uint32_t id;
    char name[88];
    double coefficient;
    uint8_t flags;
} LFQCeEfg_ctx_t;

void init_ctx(LFQCeEfg_ctx_t *ctx) {
    ctx->id = 2321;
    ctx->coefficient = 5.9;
    ctx->flags = 46;
}

int compute_VXOYHyr(int input) {
    int result = input * 7;
    for(int i = 0; i < 7; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[79];
    double coefficient;
    uint8_t flags;
} sFwRbZyY_ctx_t;

void init_ctx(sFwRbZyY_ctx_t *ctx) {
    ctx->id = 1450;
    ctx->coefficient = 8.0;
    ctx->flags = 26;
}

int compute_x8EspF5(int input) {
    int result = input * 32;
    for(int i = 0; i < 4; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[68];
    double coefficient;
    uint8_t flags;
} 2yI3mcrb_ctx_t;

void init_ctx(2yI3mcrb_ctx_t *ctx) {
    ctx->id = 9088;
    ctx->coefficient = 7.3;
    ctx->flags = 194;
}

typedef struct {
    uint32_t id;
    char name[65];
    double coefficient;
    uint8_t flags;
} NdL29L0Q_ctx_t;

void init_ctx(NdL29L0Q_ctx_t *ctx) {
    ctx->id = 2680;
    ctx->coefficient = 9.4;
    ctx->flags = 127;
}

int compute_mIaDtW1(int input) {
    int result = input * 44;
    for(int i = 0; i < 11; i++) result += (i * 4);
    return result;
}

int compute_LbZfD9d(int input) {
    int result = input * 51;
    for(int i = 0; i < 4; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[77];
    double coefficient;
    uint8_t flags;
} 9aCC6ZQT_ctx_t;

void init_ctx(9aCC6ZQT_ctx_t *ctx) {
    ctx->id = 9935;
    ctx->coefficient = 0.9;
    ctx->flags = 50;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} 4QS8haVK_ctx_t;

void init_ctx(4QS8haVK_ctx_t *ctx) {
    ctx->id = 9638;
    ctx->coefficient = 3.0;
    ctx->flags = 184;
}

typedef struct {
    uint32_t id;
    char name[61];
    double coefficient;
    uint8_t flags;
} Bt9URMod_ctx_t;

void init_ctx(Bt9URMod_ctx_t *ctx) {
    ctx->id = 8253;
    ctx->coefficient = 3.5;
    ctx->flags = 39;
}

typedef struct {
    uint32_t id;
    char name[57];
    double coefficient;
    uint8_t flags;
} 0P1JbAjM_ctx_t;

void init_ctx(0P1JbAjM_ctx_t *ctx) {
    ctx->id = 9105;
    ctx->coefficient = 1.1;
    ctx->flags = 115;
}

int compute_mtwYiSV(int input) {
    int result = input * 20;
    for(int i = 0; i < 11; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} ihwqyCry_ctx_t;

void init_ctx(ihwqyCry_ctx_t *ctx) {
    ctx->id = 2475;
    ctx->coefficient = 5.3;
    ctx->flags = 111;
}

int compute_BaGLrOw(int input) {
    int result = input * 65;
    for(int i = 0; i < 4; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} 3q1Tu3dY_ctx_t;

void init_ctx(3q1Tu3dY_ctx_t *ctx) {
    ctx->id = 5618;
    ctx->coefficient = 6.8;
    ctx->flags = 62;
}

int compute_zQaJfiG(int input) {
    int result = input * 21;
    for(int i = 0; i < 7; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[39];
    double coefficient;
    uint8_t flags;
} ut9FXscl_ctx_t;

void init_ctx(ut9FXscl_ctx_t *ctx) {
    ctx->id = 1831;
    ctx->coefficient = 2.7;
    ctx->flags = 123;
}

typedef struct {
    uint32_t id;
    char name[56];
    double coefficient;
    uint8_t flags;
} i6hjjbts_ctx_t;

void init_ctx(i6hjjbts_ctx_t *ctx) {
    ctx->id = 9999;
    ctx->coefficient = 0.4;
    ctx->flags = 31;
}

typedef struct {
    uint32_t id;
    char name[95];
    double coefficient;
    uint8_t flags;
} 7yhF8hiL_ctx_t;

void init_ctx(7yhF8hiL_ctx_t *ctx) {
    ctx->id = 9669;
    ctx->coefficient = 3.7;
    ctx->flags = 28;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} HdMnKTAs_ctx_t;

void init_ctx(HdMnKTAs_ctx_t *ctx) {
    ctx->id = 9119;
    ctx->coefficient = 2.2;
    ctx->flags = 118;
}

int compute_BmfrIfv(int input) {
    int result = input * 69;
    for(int i = 0; i < 10; i++) result += (i * 5);
    return result;
}

int compute_544n6ZF(int input) {
    int result = input * 24;
    for(int i = 0; i < 18; i++) result += (i * 3);
    return result;
}

int compute_9QWJ1gE(int input) {
    int result = input * 41;
    for(int i = 0; i < 12; i++) result += (i * 5);
    return result;
}

int compute_6ogfutb(int input) {
    int result = input * 39;
    for(int i = 0; i < 18; i++) result += (i * 4);
    return result;
}

int compute_tsyZEFI(int input) {
    int result = input * 8;
    for(int i = 0; i < 13; i++) result += (i * 4);
    return result;
}

int compute_HfUDFbK(int input) {
    int result = input * 67;
    for(int i = 0; i < 18; i++) result += (i * 1);
    return result;
}

int compute_dR82DPh(int input) {
    int result = input * 50;
    for(int i = 0; i < 7; i++) result += (i * 1);
    return result;
}

int compute_oTq9APe(int input) {
    int result = input * 47;
    for(int i = 0; i < 15; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[59];
    double coefficient;
    uint8_t flags;
} KqdKaylZ_ctx_t;

void init_ctx(KqdKaylZ_ctx_t *ctx) {
    ctx->id = 2773;
    ctx->coefficient = 8.7;
    ctx->flags = 69;
}

typedef struct {
    uint32_t id;
    char name[84];
    double coefficient;
    uint8_t flags;
} 00JYAAxJ_ctx_t;

void init_ctx(00JYAAxJ_ctx_t *ctx) {
    ctx->id = 7458;
    ctx->coefficient = 7.5;
    ctx->flags = 94;
}

int compute_Bhv5aYS(int input) {
    int result = input * 31;
    for(int i = 0; i < 8; i++) result += (i * 2);
    return result;
}

int compute_VWhNWYw(int input) {
    int result = input * 54;
    for(int i = 0; i < 17; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} KLglERkS_ctx_t;

void init_ctx(KLglERkS_ctx_t *ctx) {
    ctx->id = 4353;
    ctx->coefficient = 4.2;
    ctx->flags = 86;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} aliV4OIO_ctx_t;

void init_ctx(aliV4OIO_ctx_t *ctx) {
    ctx->id = 9463;
    ctx->coefficient = 8.7;
    ctx->flags = 73;
}

typedef struct {
    uint32_t id;
    char name[37];
    double coefficient;
    uint8_t flags;
} Yy03SHcm_ctx_t;

void init_ctx(Yy03SHcm_ctx_t *ctx) {
    ctx->id = 8826;
    ctx->coefficient = 1.7;
    ctx->flags = 234;
}

typedef struct {
    uint32_t id;
    char name[58];
    double coefficient;
    uint8_t flags;
} hZixAJZx_ctx_t;

void init_ctx(hZixAJZx_ctx_t *ctx) {
    ctx->id = 1832;
    ctx->coefficient = 3.7;
    ctx->flags = 20;
}

int compute_bbxh6Gp(int input) {
    int result = input * 45;
    for(int i = 0; i < 10; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[38];
    double coefficient;
    uint8_t flags;
} 4P02sIL3_ctx_t;

void init_ctx(4P02sIL3_ctx_t *ctx) {
    ctx->id = 5155;
    ctx->coefficient = 4.3;
    ctx->flags = 68;
}

typedef struct {
    uint32_t id;
    char name[53];
    double coefficient;
    uint8_t flags;
} L94ZUuFH_ctx_t;

void init_ctx(L94ZUuFH_ctx_t *ctx) {
    ctx->id = 8507;
    ctx->coefficient = 8.1;
    ctx->flags = 10;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} XKpBbccm_ctx_t;

void init_ctx(XKpBbccm_ctx_t *ctx) {
    ctx->id = 5420;
    ctx->coefficient = 1.3;
    ctx->flags = 3;
}

typedef struct {
    uint32_t id;
    char name[87];
    double coefficient;
    uint8_t flags;
} 5tIxMwUv_ctx_t;

void init_ctx(5tIxMwUv_ctx_t *ctx) {
    ctx->id = 3448;
    ctx->coefficient = 7.0;
    ctx->flags = 158;
}

typedef struct {
    uint32_t id;
    char name[45];
    double coefficient;
    uint8_t flags;
} 6LZqnffg_ctx_t;

void init_ctx(6LZqnffg_ctx_t *ctx) {
    ctx->id = 4601;
    ctx->coefficient = 6.8;
    ctx->flags = 92;
}

typedef struct {
    uint32_t id;
    char name[75];
    double coefficient;
    uint8_t flags;
} k9o9GJ0n_ctx_t;

void init_ctx(k9o9GJ0n_ctx_t *ctx) {
    ctx->id = 3148;
    ctx->coefficient = 1.2;
    ctx->flags = 229;
}

int compute_7kl1Zi4(int input) {
    int result = input * 65;
    for(int i = 0; i < 8; i++) result += (i * 3);
    return result;
}

int compute_k80xTkb(int input) {
    int result = input * 15;
    for(int i = 0; i < 12; i++) result += (i * 2);
    return result;
}

int compute_XExMDgj(int input) {
    int result = input * 48;
    for(int i = 0; i < 17; i++) result += (i * 2);
    return result;
}

int compute_358Trt0(int input) {
    int result = input * 14;
    for(int i = 0; i < 3; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[56];
    double coefficient;
    uint8_t flags;
} uyYTGwcM_ctx_t;

void init_ctx(uyYTGwcM_ctx_t *ctx) {
    ctx->id = 1038;
    ctx->coefficient = 2.1;
    ctx->flags = 223;
}

int compute_i4TeI9Z(int input) {
    int result = input * 40;
    for(int i = 0; i < 12; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[40];
    double coefficient;
    uint8_t flags;
} dzOzHy2H_ctx_t;

void init_ctx(dzOzHy2H_ctx_t *ctx) {
    ctx->id = 7569;
    ctx->coefficient = 6.4;
    ctx->flags = 200;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} sRLmaJSV_ctx_t;

void init_ctx(sRLmaJSV_ctx_t *ctx) {
    ctx->id = 3569;
    ctx->coefficient = 8.3;
    ctx->flags = 55;
}

int compute_EQnHmxb(int input) {
    int result = input * 67;
    for(int i = 0; i < 17; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[87];
    double coefficient;
    uint8_t flags;
} pjBqORlj_ctx_t;

void init_ctx(pjBqORlj_ctx_t *ctx) {
    ctx->id = 2095;
    ctx->coefficient = 2.0;
    ctx->flags = 221;
}

typedef struct {
    uint32_t id;
    char name[91];
    double coefficient;
    uint8_t flags;
} bmrOdNZi_ctx_t;

void init_ctx(bmrOdNZi_ctx_t *ctx) {
    ctx->id = 3603;
    ctx->coefficient = 8.8;
    ctx->flags = 193;
}

typedef struct {
    uint32_t id;
    char name[89];
    double coefficient;
    uint8_t flags;
} wM2lcd3K_ctx_t;

void init_ctx(wM2lcd3K_ctx_t *ctx) {
    ctx->id = 5760;
    ctx->coefficient = 9.4;
    ctx->flags = 102;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} FV9HWh9U_ctx_t;

void init_ctx(FV9HWh9U_ctx_t *ctx) {
    ctx->id = 5281;
    ctx->coefficient = 0.1;
    ctx->flags = 57;
}

int compute_I4mDLOm(int input) {
    int result = input * 34;
    for(int i = 0; i < 17; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[89];
    double coefficient;
    uint8_t flags;
} YTQus6DN_ctx_t;

void init_ctx(YTQus6DN_ctx_t *ctx) {
    ctx->id = 5851;
    ctx->coefficient = 0.4;
    ctx->flags = 244;
}

typedef struct {
    uint32_t id;
    char name[38];
    double coefficient;
    uint8_t flags;
} jfAPAQBl_ctx_t;

void init_ctx(jfAPAQBl_ctx_t *ctx) {
    ctx->id = 6426;
    ctx->coefficient = 6.6;
    ctx->flags = 254;
}

int compute_cXlgbhs(int input) {
    int result = input * 47;
    for(int i = 0; i < 6; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[83];
    double coefficient;
    uint8_t flags;
} MoLTlNaD_ctx_t;

void init_ctx(MoLTlNaD_ctx_t *ctx) {
    ctx->id = 8506;
    ctx->coefficient = 9.1;
    ctx->flags = 117;
}

typedef struct {
    uint32_t id;
    char name[82];
    double coefficient;
    uint8_t flags;
} Ms8XOJQ1_ctx_t;

void init_ctx(Ms8XOJQ1_ctx_t *ctx) {
    ctx->id = 6849;
    ctx->coefficient = 5.7;
    ctx->flags = 10;
}

int compute_WuV5ObZ(int input) {
    int result = input * 13;
    for(int i = 0; i < 4; i++) result += (i * 2);
    return result;
}

int compute_lRgSPUV(int input) {
    int result = input * 55;
    for(int i = 0; i < 11; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[51];
    double coefficient;
    uint8_t flags;
} ANYyu2kj_ctx_t;

void init_ctx(ANYyu2kj_ctx_t *ctx) {
    ctx->id = 4043;
    ctx->coefficient = 2.8;
    ctx->flags = 201;
}

int compute_cEveSNx(int input) {
    int result = input * 29;
    for(int i = 0; i < 6; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[37];
    double coefficient;
    uint8_t flags;
} j4eh3anI_ctx_t;

void init_ctx(j4eh3anI_ctx_t *ctx) {
    ctx->id = 6805;
    ctx->coefficient = 1.8;
    ctx->flags = 13;
}

int compute_nTBxx8l(int input) {
    int result = input * 35;
    for(int i = 0; i < 15; i++) result += (i * 4);
    return result;
}

int compute_WiEYmBj(int input) {
    int result = input * 6;
    for(int i = 0; i < 12; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[40];
    double coefficient;
    uint8_t flags;
} 7SGy1paX_ctx_t;

void init_ctx(7SGy1paX_ctx_t *ctx) {
    ctx->id = 6240;
    ctx->coefficient = 3.3;
    ctx->flags = 255;
}

int compute_a3uaj93(int input) {
    int result = input * 57;
    for(int i = 0; i < 16; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[77];
    double coefficient;
    uint8_t flags;
} I94qyOPl_ctx_t;

void init_ctx(I94qyOPl_ctx_t *ctx) {
    ctx->id = 4598;
    ctx->coefficient = 2.2;
    ctx->flags = 126;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} gGu1ZxIu_ctx_t;

void init_ctx(gGu1ZxIu_ctx_t *ctx) {
    ctx->id = 5896;
    ctx->coefficient = 9.1;
    ctx->flags = 25;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} cxQhndiB_ctx_t;

void init_ctx(cxQhndiB_ctx_t *ctx) {
    ctx->id = 4838;
    ctx->coefficient = 6.4;
    ctx->flags = 146;
}

typedef struct {
    uint32_t id;
    char name[93];
    double coefficient;
    uint8_t flags;
} zW9QBash_ctx_t;

void init_ctx(zW9QBash_ctx_t *ctx) {
    ctx->id = 8586;
    ctx->coefficient = 5.1;
    ctx->flags = 97;
}

typedef struct {
    uint32_t id;
    char name[88];
    double coefficient;
    uint8_t flags;
} aPlCm6Z2_ctx_t;

void init_ctx(aPlCm6Z2_ctx_t *ctx) {
    ctx->id = 5764;
    ctx->coefficient = 2.1;
    ctx->flags = 32;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} jKRWM7vk_ctx_t;

void init_ctx(jKRWM7vk_ctx_t *ctx) {
    ctx->id = 2181;
    ctx->coefficient = 8.5;
    ctx->flags = 86;
}

typedef struct {
    uint32_t id;
    char name[33];
    double coefficient;
    uint8_t flags;
} NWORuAIT_ctx_t;

void init_ctx(NWORuAIT_ctx_t *ctx) {
    ctx->id = 6116;
    ctx->coefficient = 1.2;
    ctx->flags = 72;
}

typedef struct {
    uint32_t id;
    char name[89];
    double coefficient;
    uint8_t flags;
} qTlMminu_ctx_t;

void init_ctx(qTlMminu_ctx_t *ctx) {
    ctx->id = 5034;
    ctx->coefficient = 8.6;
    ctx->flags = 36;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} rkqubeBS_ctx_t;

void init_ctx(rkqubeBS_ctx_t *ctx) {
    ctx->id = 7957;
    ctx->coefficient = 0.3;
    ctx->flags = 232;
}

int compute_FAiPx8l(int input) {
    int result = input * 58;
    for(int i = 0; i < 14; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} r8tzHZm1_ctx_t;

void init_ctx(r8tzHZm1_ctx_t *ctx) {
    ctx->id = 8352;
    ctx->coefficient = 3.2;
    ctx->flags = 134;
}

typedef struct {
    uint32_t id;
    char name[90];
    double coefficient;
    uint8_t flags;
} yZxNLW9u_ctx_t;

void init_ctx(yZxNLW9u_ctx_t *ctx) {
    ctx->id = 7425;
    ctx->coefficient = 7.9;
    ctx->flags = 50;
}

typedef struct {
    uint32_t id;
    char name[94];
    double coefficient;
    uint8_t flags;
} vzwyTF1K_ctx_t;

void init_ctx(vzwyTF1K_ctx_t *ctx) {
    ctx->id = 8068;
    ctx->coefficient = 6.4;
    ctx->flags = 97;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} UJ2iOCdR_ctx_t;

void init_ctx(UJ2iOCdR_ctx_t *ctx) {
    ctx->id = 8411;
    ctx->coefficient = 5.1;
    ctx->flags = 240;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} 9Qms2HiU_ctx_t;

void init_ctx(9Qms2HiU_ctx_t *ctx) {
    ctx->id = 4768;
    ctx->coefficient = 2.8;
    ctx->flags = 58;
}

typedef struct {
    uint32_t id;
    char name[40];
    double coefficient;
    uint8_t flags;
} ktRNZAhw_ctx_t;

void init_ctx(ktRNZAhw_ctx_t *ctx) {
    ctx->id = 8660;
    ctx->coefficient = 7.1;
    ctx->flags = 131;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} LGBtXYbK_ctx_t;

void init_ctx(LGBtXYbK_ctx_t *ctx) {
    ctx->id = 1699;
    ctx->coefficient = 4.4;
    ctx->flags = 245;
}

int compute_Fx4oTcZ(int input) {
    int result = input * 65;
    for(int i = 0; i < 15; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[58];
    double coefficient;
    uint8_t flags;
} TkSUP4Bx_ctx_t;

void init_ctx(TkSUP4Bx_ctx_t *ctx) {
    ctx->id = 9561;
    ctx->coefficient = 6.3;
    ctx->flags = 7;
}

typedef struct {
    uint32_t id;
    char name[63];
    double coefficient;
    uint8_t flags;
} bcRezxqW_ctx_t;

void init_ctx(bcRezxqW_ctx_t *ctx) {
    ctx->id = 5861;
    ctx->coefficient = 6.7;
    ctx->flags = 102;
}

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} YxxTFX1L_ctx_t;

void init_ctx(YxxTFX1L_ctx_t *ctx) {
    ctx->id = 3721;
    ctx->coefficient = 7.1;
    ctx->flags = 245;
}

typedef struct {
    uint32_t id;
    char name[52];
    double coefficient;
    uint8_t flags;
} Y2YuCc9o_ctx_t;

void init_ctx(Y2YuCc9o_ctx_t *ctx) {
    ctx->id = 7311;
    ctx->coefficient = 5.1;
    ctx->flags = 49;
}

int compute_sAIGmup(int input) {
    int result = input * 9;
    for(int i = 0; i < 7; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[38];
    double coefficient;
    uint8_t flags;
} T1hHyKgD_ctx_t;

void init_ctx(T1hHyKgD_ctx_t *ctx) {
    ctx->id = 7070;
    ctx->coefficient = 3.7;
    ctx->flags = 29;
}

typedef struct {
    uint32_t id;
    char name[72];
    double coefficient;
    uint8_t flags;
} HWmeORsW_ctx_t;

void init_ctx(HWmeORsW_ctx_t *ctx) {
    ctx->id = 6447;
    ctx->coefficient = 8.5;
    ctx->flags = 44;
}

typedef struct {
    uint32_t id;
    char name[37];
    double coefficient;
    uint8_t flags;
} 3XS77JeD_ctx_t;

void init_ctx(3XS77JeD_ctx_t *ctx) {
    ctx->id = 6952;
    ctx->coefficient = 4.8;
    ctx->flags = 18;
}

typedef struct {
    uint32_t id;
    char name[79];
    double coefficient;
    uint8_t flags;
} 4dDJ13T7_ctx_t;

void init_ctx(4dDJ13T7_ctx_t *ctx) {
    ctx->id = 5745;
    ctx->coefficient = 5.4;
    ctx->flags = 88;
}

int compute_eTrocfZ(int input) {
    int result = input * 43;
    for(int i = 0; i < 12; i++) result += (i * 3);
    return result;
}

int compute_9jEvafU(int input) {
    int result = input * 65;
    for(int i = 0; i < 10; i++) result += (i * 3);
    return result;
}

int compute_L9Innlu(int input) {
    int result = input * 20;
    for(int i = 0; i < 15; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[63];
    double coefficient;
    uint8_t flags;
} 7m6sSbNX_ctx_t;

void init_ctx(7m6sSbNX_ctx_t *ctx) {
    ctx->id = 7096;
    ctx->coefficient = 8.5;
    ctx->flags = 160;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} lU5ffDnT_ctx_t;

void init_ctx(lU5ffDnT_ctx_t *ctx) {
    ctx->id = 5764;
    ctx->coefficient = 6.1;
    ctx->flags = 78;
}

typedef struct {
    uint32_t id;
    char name[62];
    double coefficient;
    uint8_t flags;
} rIoMXUR8_ctx_t;

void init_ctx(rIoMXUR8_ctx_t *ctx) {
    ctx->id = 6689;
    ctx->coefficient = 3.4;
    ctx->flags = 76;
}

int compute_uzpNCfi(int input) {
    int result = input * 9;
    for(int i = 0; i < 18; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[37];
    double coefficient;
    uint8_t flags;
} 8OArzFsB_ctx_t;

void init_ctx(8OArzFsB_ctx_t *ctx) {
    ctx->id = 8680;
    ctx->coefficient = 5.0;
    ctx->flags = 217;
}

typedef struct {
    uint32_t id;
    char name[87];
    double coefficient;
    uint8_t flags;
} hLYyCvs7_ctx_t;

void init_ctx(hLYyCvs7_ctx_t *ctx) {
    ctx->id = 8468;
    ctx->coefficient = 9.5;
    ctx->flags = 34;
}

typedef struct {
    uint32_t id;
    char name[65];
    double coefficient;
    uint8_t flags;
} fothHhWp_ctx_t;

void init_ctx(fothHhWp_ctx_t *ctx) {
    ctx->id = 1725;
    ctx->coefficient = 8.2;
    ctx->flags = 105;
}

typedef struct {
    uint32_t id;
    char name[76];
    double coefficient;
    uint8_t flags;
} TKz7RPPj_ctx_t;

void init_ctx(TKz7RPPj_ctx_t *ctx) {
    ctx->id = 1734;
    ctx->coefficient = 8.9;
    ctx->flags = 15;
}

typedef struct {
    uint32_t id;
    char name[73];
    double coefficient;
    uint8_t flags;
} bbv4G2Dx_ctx_t;

void init_ctx(bbv4G2Dx_ctx_t *ctx) {
    ctx->id = 8246;
    ctx->coefficient = 8.2;
    ctx->flags = 239;
}

typedef struct {
    uint32_t id;
    char name[35];
    double coefficient;
    uint8_t flags;
} k5umaEvk_ctx_t;

void init_ctx(k5umaEvk_ctx_t *ctx) {
    ctx->id = 9336;
    ctx->coefficient = 4.5;
    ctx->flags = 3;
}

int compute_zSsjmLw(int input) {
    int result = input * 27;
    for(int i = 0; i < 6; i++) result += (i * 2);
    return result;
}

int compute_zCaRLFv(int input) {
    int result = input * 69;
    for(int i = 0; i < 5; i++) result += (i * 1);
    return result;
}

int compute_C4eSCr9(int input) {
    int result = input * 69;
    for(int i = 0; i < 10; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[84];
    double coefficient;
    uint8_t flags;
} rXC8zG32_ctx_t;

void init_ctx(rXC8zG32_ctx_t *ctx) {
    ctx->id = 8982;
    ctx->coefficient = 0.4;
    ctx->flags = 27;
}

typedef struct {
    uint32_t id;
    char name[68];
    double coefficient;
    uint8_t flags;
} FDkrVPjm_ctx_t;

void init_ctx(FDkrVPjm_ctx_t *ctx) {
    ctx->id = 9368;
    ctx->coefficient = 9.4;
    ctx->flags = 215;
}

int compute_iLpu5Pg(int input) {
    int result = input * 6;
    for(int i = 0; i < 3; i++) result += (i * 1);
    return result;
}

int compute_o2bL1Xb(int input) {
    int result = input * 24;
    for(int i = 0; i < 5; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[76];
    double coefficient;
    uint8_t flags;
} jpIViN7T_ctx_t;

void init_ctx(jpIViN7T_ctx_t *ctx) {
    ctx->id = 7380;
    ctx->coefficient = 5.4;
    ctx->flags = 18;
}

typedef struct {
    uint32_t id;
    char name[39];
    double coefficient;
    uint8_t flags;
} bcwuP6MO_ctx_t;

void init_ctx(bcwuP6MO_ctx_t *ctx) {
    ctx->id = 2555;
    ctx->coefficient = 8.5;
    ctx->flags = 219;
}

typedef struct {
    uint32_t id;
    char name[63];
    double coefficient;
    uint8_t flags;
} aBzXyWQS_ctx_t;

void init_ctx(aBzXyWQS_ctx_t *ctx) {
    ctx->id = 3555;
    ctx->coefficient = 2.9;
    ctx->flags = 81;
}

int compute_5NDca8Q(int input) {
    int result = input * 25;
    for(int i = 0; i < 14; i++) result += (i * 4);
    return result;
}

int compute_Ntazrv6(int input) {
    int result = input * 69;
    for(int i = 0; i < 13; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[39];
    double coefficient;
    uint8_t flags;
} 2HRKFXNz_ctx_t;

void init_ctx(2HRKFXNz_ctx_t *ctx) {
    ctx->id = 7916;
    ctx->coefficient = 2.3;
    ctx->flags = 95;
}

int compute_AkAbucs(int input) {
    int result = input * 37;
    for(int i = 0; i < 13; i++) result += (i * 3);
    return result;
}

int compute_2VrOQ5U(int input) {
    int result = input * 69;
    for(int i = 0; i < 3; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[92];
    double coefficient;
    uint8_t flags;
} 8fnZ7y7M_ctx_t;

void init_ctx(8fnZ7y7M_ctx_t *ctx) {
    ctx->id = 8538;
    ctx->coefficient = 7.5;
    ctx->flags = 210;
}

typedef struct {
    uint32_t id;
    char name[63];
    double coefficient;
    uint8_t flags;
} QDeo9yXS_ctx_t;

void init_ctx(QDeo9yXS_ctx_t *ctx) {
    ctx->id = 7194;
    ctx->coefficient = 8.2;
    ctx->flags = 137;
}

typedef struct {
    uint32_t id;
    char name[74];
    double coefficient;
    uint8_t flags;
} jj1fOO3h_ctx_t;

void init_ctx(jj1fOO3h_ctx_t *ctx) {
    ctx->id = 7952;
    ctx->coefficient = 8.8;
    ctx->flags = 226;
}

typedef struct {
    uint32_t id;
    char name[72];
    double coefficient;
    uint8_t flags;
} QPM1Af0e_ctx_t;

void init_ctx(QPM1Af0e_ctx_t *ctx) {
    ctx->id = 1839;
    ctx->coefficient = 5.4;
    ctx->flags = 78;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} gZxCNCWo_ctx_t;

void init_ctx(gZxCNCWo_ctx_t *ctx) {
    ctx->id = 9376;
    ctx->coefficient = 9.1;
    ctx->flags = 113;
}

int compute_WyEelas(int input) {
    int result = input * 17;
    for(int i = 0; i < 5; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[80];
    double coefficient;
    uint8_t flags;
} gE4oMdmO_ctx_t;

void init_ctx(gE4oMdmO_ctx_t *ctx) {
    ctx->id = 5212;
    ctx->coefficient = 4.8;
    ctx->flags = 135;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} S4XYvTRE_ctx_t;

void init_ctx(S4XYvTRE_ctx_t *ctx) {
    ctx->id = 4012;
    ctx->coefficient = 5.4;
    ctx->flags = 3;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} gHNswMV9_ctx_t;

void init_ctx(gHNswMV9_ctx_t *ctx) {
    ctx->id = 1826;
    ctx->coefficient = 9.3;
    ctx->flags = 119;
}

int compute_psAhU2P(int input) {
    int result = input * 65;
    for(int i = 0; i < 14; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[95];
    double coefficient;
    uint8_t flags;
} r481Hps7_ctx_t;

void init_ctx(r481Hps7_ctx_t *ctx) {
    ctx->id = 3703;
    ctx->coefficient = 8.5;
    ctx->flags = 243;
}

typedef struct {
    uint32_t id;
    char name[91];
    double coefficient;
    uint8_t flags;
} ThreZVcH_ctx_t;

void init_ctx(ThreZVcH_ctx_t *ctx) {
    ctx->id = 5903;
    ctx->coefficient = 1.6;
    ctx->flags = 157;
}

typedef struct {
    uint32_t id;
    char name[63];
    double coefficient;
    uint8_t flags;
} emiEgalS_ctx_t;

void init_ctx(emiEgalS_ctx_t *ctx) {
    ctx->id = 3449;
    ctx->coefficient = 2.3;
    ctx->flags = 94;
}

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} NX0u3sBL_ctx_t;

void init_ctx(NX0u3sBL_ctx_t *ctx) {
    ctx->id = 7619;
    ctx->coefficient = 4.4;
    ctx->flags = 154;
}

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} hqHpZrJf_ctx_t;

void init_ctx(hqHpZrJf_ctx_t *ctx) {
    ctx->id = 7811;
    ctx->coefficient = 3.0;
    ctx->flags = 64;
}

typedef struct {
    uint32_t id;
    char name[59];
    double coefficient;
    uint8_t flags;
} wR0zrwGD_ctx_t;

void init_ctx(wR0zrwGD_ctx_t *ctx) {
    ctx->id = 2401;
    ctx->coefficient = 5.2;
    ctx->flags = 129;
}

int compute_ROtHysS(int input) {
    int result = input * 4;
    for(int i = 0; i < 17; i++) result += (i * 5);
    return result;
}

int compute_mubannG(int input) {
    int result = input * 47;
    for(int i = 0; i < 15; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[94];
    double coefficient;
    uint8_t flags;
} bG9zBjiQ_ctx_t;

void init_ctx(bG9zBjiQ_ctx_t *ctx) {
    ctx->id = 6992;
    ctx->coefficient = 0.8;
    ctx->flags = 10;
}

typedef struct {
    uint32_t id;
    char name[80];
    double coefficient;
    uint8_t flags;
} NOz9RMQw_ctx_t;

void init_ctx(NOz9RMQw_ctx_t *ctx) {
    ctx->id = 6563;
    ctx->coefficient = 3.2;
    ctx->flags = 82;
}

typedef struct {
    uint32_t id;
    char name[90];
    double coefficient;
    uint8_t flags;
} cWqCjFj2_ctx_t;

void init_ctx(cWqCjFj2_ctx_t *ctx) {
    ctx->id = 9031;
    ctx->coefficient = 7.9;
    ctx->flags = 106;
}

typedef struct {
    uint32_t id;
    char name[75];
    double coefficient;
    uint8_t flags;
} IYcGGWBR_ctx_t;

void init_ctx(IYcGGWBR_ctx_t *ctx) {
    ctx->id = 6132;
    ctx->coefficient = 8.3;
    ctx->flags = 98;
}

typedef struct {
    uint32_t id;
    char name[82];
    double coefficient;
    uint8_t flags;
} bLrxl8uD_ctx_t;

void init_ctx(bLrxl8uD_ctx_t *ctx) {
    ctx->id = 1486;
    ctx->coefficient = 3.1;
    ctx->flags = 188;
}

typedef struct {
    uint32_t id;
    char name[76];
    double coefficient;
    uint8_t flags;
} BdXKDOPE_ctx_t;

void init_ctx(BdXKDOPE_ctx_t *ctx) {
    ctx->id = 8525;
    ctx->coefficient = 8.8;
    ctx->flags = 132;
}

typedef struct {
    uint32_t id;
    char name[67];
    double coefficient;
    uint8_t flags;
} w5JKpH5U_ctx_t;

void init_ctx(w5JKpH5U_ctx_t *ctx) {
    ctx->id = 8680;
    ctx->coefficient = 5.5;
    ctx->flags = 12;
}

int compute_WAH5FAH(int input) {
    int result = input * 29;
    for(int i = 0; i < 3; i++) result += (i * 4);
    return result;
}

int compute_EUsZddr(int input) {
    int result = input * 20;
    for(int i = 0; i < 8; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[59];
    double coefficient;
    uint8_t flags;
} WZMI9WBU_ctx_t;

void init_ctx(WZMI9WBU_ctx_t *ctx) {
    ctx->id = 4611;
    ctx->coefficient = 1.7;
    ctx->flags = 127;
}

typedef struct {
    uint32_t id;
    char name[62];
    double coefficient;
    uint8_t flags;
} eKmgss0j_ctx_t;

void init_ctx(eKmgss0j_ctx_t *ctx) {
    ctx->id = 3577;
    ctx->coefficient = 3.5;
    ctx->flags = 231;
}

typedef struct {
    uint32_t id;
    char name[84];
    double coefficient;
    uint8_t flags;
} wupLbjEX_ctx_t;

void init_ctx(wupLbjEX_ctx_t *ctx) {
    ctx->id = 1007;
    ctx->coefficient = 1.3;
    ctx->flags = 169;
}

int compute_8PixKCc(int input) {
    int result = input * 58;
    for(int i = 0; i < 8; i++) result += (i * 2);
    return result;
}

int compute_Nnsa7dB(int input) {
    int result = input * 35;
    for(int i = 0; i < 18; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} GerrYZuS_ctx_t;

void init_ctx(GerrYZuS_ctx_t *ctx) {
    ctx->id = 2996;
    ctx->coefficient = 9.5;
    ctx->flags = 25;
}

typedef struct {
    uint32_t id;
    char name[92];
    double coefficient;
    uint8_t flags;
} Z0YGY7Je_ctx_t;

void init_ctx(Z0YGY7Je_ctx_t *ctx) {
    ctx->id = 5766;
    ctx->coefficient = 5.6;
    ctx->flags = 64;
}

int compute_nVJZcVJ(int input) {
    int result = input * 66;
    for(int i = 0; i < 9; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[96];
    double coefficient;
    uint8_t flags;
} yD5zbmFF_ctx_t;

void init_ctx(yD5zbmFF_ctx_t *ctx) {
    ctx->id = 2924;
    ctx->coefficient = 10.0;
    ctx->flags = 233;
}

typedef struct {
    uint32_t id;
    char name[47];
    double coefficient;
    uint8_t flags;
} HnlAsC73_ctx_t;

void init_ctx(HnlAsC73_ctx_t *ctx) {
    ctx->id = 4888;
    ctx->coefficient = 6.9;
    ctx->flags = 221;
}

int compute_bVDowG1(int input) {
    int result = input * 25;
    for(int i = 0; i < 15; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} hH5hmWRs_ctx_t;

void init_ctx(hH5hmWRs_ctx_t *ctx) {
    ctx->id = 1763;
    ctx->coefficient = 9.0;
    ctx->flags = 173;
}

int compute_Ol753Dx(int input) {
    int result = input * 69;
    for(int i = 0; i < 9; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} LgZPiFl8_ctx_t;

void init_ctx(LgZPiFl8_ctx_t *ctx) {
    ctx->id = 8441;
    ctx->coefficient = 9.6;
    ctx->flags = 87;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} CAWA3vvX_ctx_t;

void init_ctx(CAWA3vvX_ctx_t *ctx) {
    ctx->id = 5330;
    ctx->coefficient = 5.7;
    ctx->flags = 145;
}

typedef struct {
    uint32_t id;
    char name[68];
    double coefficient;
    uint8_t flags;
} wX4AXorl_ctx_t;

void init_ctx(wX4AXorl_ctx_t *ctx) {
    ctx->id = 8423;
    ctx->coefficient = 1.6;
    ctx->flags = 150;
}

typedef struct {
    uint32_t id;
    char name[38];
    double coefficient;
    uint8_t flags;
} Xi17PR29_ctx_t;

void init_ctx(Xi17PR29_ctx_t *ctx) {
    ctx->id = 8906;
    ctx->coefficient = 6.8;
    ctx->flags = 30;
}

typedef struct {
    uint32_t id;
    char name[51];
    double coefficient;
    uint8_t flags;
} zARkEDCE_ctx_t;

void init_ctx(zARkEDCE_ctx_t *ctx) {
    ctx->id = 2501;
    ctx->coefficient = 6.4;
    ctx->flags = 88;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} 5VzFpLwn_ctx_t;

void init_ctx(5VzFpLwn_ctx_t *ctx) {
    ctx->id = 6750;
    ctx->coefficient = 8.4;
    ctx->flags = 248;
}

int compute_cX6d3S6(int input) {
    int result = input * 54;
    for(int i = 0; i < 16; i++) result += (i * 2);
    return result;
}

int compute_8mymzQF(int input) {
    int result = input * 62;
    for(int i = 0; i < 15; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[66];
    double coefficient;
    uint8_t flags;
} qrPZKncD_ctx_t;

void init_ctx(qrPZKncD_ctx_t *ctx) {
    ctx->id = 7553;
    ctx->coefficient = 7.8;
    ctx->flags = 103;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} c5Vphscm_ctx_t;

void init_ctx(c5Vphscm_ctx_t *ctx) {
    ctx->id = 5075;
    ctx->coefficient = 5.0;
    ctx->flags = 223;
}

int compute_VHW9zGM(int input) {
    int result = input * 11;
    for(int i = 0; i < 3; i++) result += (i * 1);
    return result;
}

int compute_8hLAY09(int input) {
    int result = input * 40;
    for(int i = 0; i < 18; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[77];
    double coefficient;
    uint8_t flags;
} AXlOaVd5_ctx_t;

void init_ctx(AXlOaVd5_ctx_t *ctx) {
    ctx->id = 6499;
    ctx->coefficient = 1.8;
    ctx->flags = 188;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} MW2QVI8J_ctx_t;

void init_ctx(MW2QVI8J_ctx_t *ctx) {
    ctx->id = 5150;
    ctx->coefficient = 5.9;
    ctx->flags = 201;
}

int compute_Q3DZFWx(int input) {
    int result = input * 41;
    for(int i = 0; i < 11; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} D5CWGdJR_ctx_t;

void init_ctx(D5CWGdJR_ctx_t *ctx) {
    ctx->id = 8255;
    ctx->coefficient = 9.0;
    ctx->flags = 75;
}

int compute_GbvMQU9(int input) {
    int result = input * 17;
    for(int i = 0; i < 14; i++) result += (i * 1);
    return result;
}

int compute_duUaldK(int input) {
    int result = input * 43;
    for(int i = 0; i < 11; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[52];
    double coefficient;
    uint8_t flags;
} c6FMi6UN_ctx_t;

void init_ctx(c6FMi6UN_ctx_t *ctx) {
    ctx->id = 7769;
    ctx->coefficient = 1.5;
    ctx->flags = 163;
}

typedef struct {
    uint32_t id;
    char name[88];
    double coefficient;
    uint8_t flags;
} nNwJOJv0_ctx_t;

void init_ctx(nNwJOJv0_ctx_t *ctx) {
    ctx->id = 6350;
    ctx->coefficient = 5.9;
    ctx->flags = 8;
}

int compute_tu2ueIq(int input) {
    int result = input * 44;
    for(int i = 0; i < 6; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[84];
    double coefficient;
    uint8_t flags;
} 4siulXas_ctx_t;

void init_ctx(4siulXas_ctx_t *ctx) {
    ctx->id = 7996;
    ctx->coefficient = 3.6;
    ctx->flags = 8;
}

typedef struct {
    uint32_t id;
    char name[33];
    double coefficient;
    uint8_t flags;
} D8lrc2zU_ctx_t;

void init_ctx(D8lrc2zU_ctx_t *ctx) {
    ctx->id = 8204;
    ctx->coefficient = 5.6;
    ctx->flags = 128;
}

typedef struct {
    uint32_t id;
    char name[88];
    double coefficient;
    uint8_t flags;
} Ud3VaT1x_ctx_t;

void init_ctx(Ud3VaT1x_ctx_t *ctx) {
    ctx->id = 8538;
    ctx->coefficient = 0.9;
    ctx->flags = 55;
}

int compute_ZnRIFWv(int input) {
    int result = input * 18;
    for(int i = 0; i < 11; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[95];
    double coefficient;
    uint8_t flags;
} snvl9QLC_ctx_t;

void init_ctx(snvl9QLC_ctx_t *ctx) {
    ctx->id = 5473;
    ctx->coefficient = 2.8;
    ctx->flags = 48;
}

typedef struct {
    uint32_t id;
    char name[66];
    double coefficient;
    uint8_t flags;
} jYQI99ed_ctx_t;

void init_ctx(jYQI99ed_ctx_t *ctx) {
    ctx->id = 9406;
    ctx->coefficient = 8.4;
    ctx->flags = 117;
}

typedef struct {
    uint32_t id;
    char name[48];
    double coefficient;
    uint8_t flags;
} e5mI1dIY_ctx_t;

void init_ctx(e5mI1dIY_ctx_t *ctx) {
    ctx->id = 1121;
    ctx->coefficient = 2.7;
    ctx->flags = 62;
}

typedef struct {
    uint32_t id;
    char name[56];
    double coefficient;
    uint8_t flags;
} fAtQNFa3_ctx_t;

void init_ctx(fAtQNFa3_ctx_t *ctx) {
    ctx->id = 7145;
    ctx->coefficient = 9.6;
    ctx->flags = 154;
}

typedef struct {
    uint32_t id;
    char name[54];
    double coefficient;
    uint8_t flags;
} A7SuqugA_ctx_t;

void init_ctx(A7SuqugA_ctx_t *ctx) {
    ctx->id = 2062;
    ctx->coefficient = 8.0;
    ctx->flags = 184;
}

typedef struct {
    uint32_t id;
    char name[53];
    double coefficient;
    uint8_t flags;
} sQWheuQU_ctx_t;

void init_ctx(sQWheuQU_ctx_t *ctx) {
    ctx->id = 3426;
    ctx->coefficient = 8.3;
    ctx->flags = 158;
}

int compute_lhe72FK(int input) {
    int result = input * 34;
    for(int i = 0; i < 16; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[82];
    double coefficient;
    uint8_t flags;
} 0qFrOn0G_ctx_t;

void init_ctx(0qFrOn0G_ctx_t *ctx) {
    ctx->id = 3457;
    ctx->coefficient = 2.0;
    ctx->flags = 119;
}

typedef struct {
    uint32_t id;
    char name[73];
    double coefficient;
    uint8_t flags;
} q5FlqpDy_ctx_t;

void init_ctx(q5FlqpDy_ctx_t *ctx) {
    ctx->id = 2600;
    ctx->coefficient = 7.7;
    ctx->flags = 2;
}

int compute_NukQPOA(int input) {
    int result = input * 20;
    for(int i = 0; i < 18; i++) result += (i * 4);
    return result;
}

int compute_BHyL5gd(int input) {
    int result = input * 63;
    for(int i = 0; i < 18; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[77];
    double coefficient;
    uint8_t flags;
} zMaU4T5n_ctx_t;

void init_ctx(zMaU4T5n_ctx_t *ctx) {
    ctx->id = 5185;
    ctx->coefficient = 1.6;
    ctx->flags = 89;
}

int compute_rHw2Sfn(int input) {
    int result = input * 7;
    for(int i = 0; i < 16; i++) result += (i * 2);
    return result;
}

int compute_lLeiNul(int input) {
    int result = input * 66;
    for(int i = 0; i < 6; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} afcGRGjr_ctx_t;

void init_ctx(afcGRGjr_ctx_t *ctx) {
    ctx->id = 9083;
    ctx->coefficient = 5.3;
    ctx->flags = 95;
}

int compute_tN3m9RQ(int input) {
    int result = input * 43;
    for(int i = 0; i < 3; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[74];
    double coefficient;
    uint8_t flags;
} ZXxYZ9C7_ctx_t;

void init_ctx(ZXxYZ9C7_ctx_t *ctx) {
    ctx->id = 4297;
    ctx->coefficient = 8.8;
    ctx->flags = 85;
}

int compute_4NbjEit(int input) {
    int result = input * 8;
    for(int i = 0; i < 12; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[78];
    double coefficient;
    uint8_t flags;
} Xy0MSni6_ctx_t;

void init_ctx(Xy0MSni6_ctx_t *ctx) {
    ctx->id = 8497;
    ctx->coefficient = 0.6;
    ctx->flags = 65;
}

int compute_xNhYcsA(int input) {
    int result = input * 33;
    for(int i = 0; i < 11; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[37];
    double coefficient;
    uint8_t flags;
} YXjKfLUu_ctx_t;

void init_ctx(YXjKfLUu_ctx_t *ctx) {
    ctx->id = 6090;
    ctx->coefficient = 8.3;
    ctx->flags = 38;
}

typedef struct {
    uint32_t id;
    char name[36];
    double coefficient;
    uint8_t flags;
} xXwRbxDo_ctx_t;

void init_ctx(xXwRbxDo_ctx_t *ctx) {
    ctx->id = 4853;
    ctx->coefficient = 3.1;
    ctx->flags = 177;
}

int compute_Gqr1Dqu(int input) {
    int result = input * 32;
    for(int i = 0; i < 12; i++) result += (i * 3);
    return result;
}

int compute_Azy7OBv(int input) {
    int result = input * 30;
    for(int i = 0; i < 8; i++) result += (i * 1);
    return result;
}

int compute_BkUL2O8(int input) {
    int result = input * 33;
    for(int i = 0; i < 16; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} rSh1DiwD_ctx_t;

void init_ctx(rSh1DiwD_ctx_t *ctx) {
    ctx->id = 7833;
    ctx->coefficient = 6.9;
    ctx->flags = 67;
}

typedef struct {
    uint32_t id;
    char name[85];
    double coefficient;
    uint8_t flags;
} u0xEzNdf_ctx_t;

void init_ctx(u0xEzNdf_ctx_t *ctx) {
    ctx->id = 8546;
    ctx->coefficient = 0.1;
    ctx->flags = 74;
}

typedef struct {
    uint32_t id;
    char name[52];
    double coefficient;
    uint8_t flags;
} lLJE0nvY_ctx_t;

void init_ctx(lLJE0nvY_ctx_t *ctx) {
    ctx->id = 2693;
    ctx->coefficient = 9.7;
    ctx->flags = 204;
}

int compute_pZ2hskS(int input) {
    int result = input * 28;
    for(int i = 0; i < 15; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[83];
    double coefficient;
    uint8_t flags;
} TZSvPdqc_ctx_t;

void init_ctx(TZSvPdqc_ctx_t *ctx) {
    ctx->id = 1712;
    ctx->coefficient = 0.8;
    ctx->flags = 224;
}

int compute_wDWBDCL(int input) {
    int result = input * 19;
    for(int i = 0; i < 13; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[49];
    double coefficient;
    uint8_t flags;
} HSwVNOKw_ctx_t;

void init_ctx(HSwVNOKw_ctx_t *ctx) {
    ctx->id = 9712;
    ctx->coefficient = 7.4;
    ctx->flags = 166;
}

typedef struct {
    uint32_t id;
    char name[69];
    double coefficient;
    uint8_t flags;
} ihN8JvSp_ctx_t;

void init_ctx(ihN8JvSp_ctx_t *ctx) {
    ctx->id = 4274;
    ctx->coefficient = 8.2;
    ctx->flags = 37;
}

int compute_eFo41Bk(int input) {
    int result = input * 52;
    for(int i = 0; i < 15; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[53];
    double coefficient;
    uint8_t flags;
} hm4O18ic_ctx_t;

void init_ctx(hm4O18ic_ctx_t *ctx) {
    ctx->id = 1550;
    ctx->coefficient = 9.4;
    ctx->flags = 50;
}

int compute_a8EaQpJ(int input) {
    int result = input * 46;
    for(int i = 0; i < 11; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[75];
    double coefficient;
    uint8_t flags;
} qwodSoLf_ctx_t;

void init_ctx(qwodSoLf_ctx_t *ctx) {
    ctx->id = 5713;
    ctx->coefficient = 9.1;
    ctx->flags = 52;
}

int compute_lp1qxei(int input) {
    int result = input * 53;
    for(int i = 0; i < 7; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[43];
    double coefficient;
    uint8_t flags;
} 3riJDNKa_ctx_t;

void init_ctx(3riJDNKa_ctx_t *ctx) {
    ctx->id = 1588;
    ctx->coefficient = 9.0;
    ctx->flags = 73;
}

typedef struct {
    uint32_t id;
    char name[44];
    double coefficient;
    uint8_t flags;
} vq6oBiae_ctx_t;

void init_ctx(vq6oBiae_ctx_t *ctx) {
    ctx->id = 1513;
    ctx->coefficient = 0.6;
    ctx->flags = 20;
}

int compute_KNss5ej(int input) {
    int result = input * 41;
    for(int i = 0; i < 11; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[95];
    double coefficient;
    uint8_t flags;
} gXG62cXx_ctx_t;

void init_ctx(gXG62cXx_ctx_t *ctx) {
    ctx->id = 3492;
    ctx->coefficient = 2.0;
    ctx->flags = 1;
}

int compute_8FleA6h(int input) {
    int result = input * 51;
    for(int i = 0; i < 12; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[96];
    double coefficient;
    uint8_t flags;
} PaBkuIgC_ctx_t;

void init_ctx(PaBkuIgC_ctx_t *ctx) {
    ctx->id = 2339;
    ctx->coefficient = 2.0;
    ctx->flags = 6;
}

typedef struct {
    uint32_t id;
    char name[46];
    double coefficient;
    uint8_t flags;
} 46hcRKMq_ctx_t;

void init_ctx(46hcRKMq_ctx_t *ctx) {
    ctx->id = 5972;
    ctx->coefficient = 5.3;
    ctx->flags = 241;
}

int compute_rXElOcU(int input) {
    int result = input * 18;
    for(int i = 0; i < 6; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[49];
    double coefficient;
    uint8_t flags;
} pDtMssnw_ctx_t;

void init_ctx(pDtMssnw_ctx_t *ctx) {
    ctx->id = 5959;
    ctx->coefficient = 9.0;
    ctx->flags = 88;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} 7FStg8xG_ctx_t;

void init_ctx(7FStg8xG_ctx_t *ctx) {
    ctx->id = 5321;
    ctx->coefficient = 0.5;
    ctx->flags = 22;
}

int compute_vdu44WP(int input) {
    int result = input * 25;
    for(int i = 0; i < 6; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[80];
    double coefficient;
    uint8_t flags;
} 96MhdeJe_ctx_t;

void init_ctx(96MhdeJe_ctx_t *ctx) {
    ctx->id = 8835;
    ctx->coefficient = 6.6;
    ctx->flags = 194;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} h6lrEn3d_ctx_t;

void init_ctx(h6lrEn3d_ctx_t *ctx) {
    ctx->id = 8198;
    ctx->coefficient = 0.7;
    ctx->flags = 219;
}

typedef struct {
    uint32_t id;
    char name[57];
    double coefficient;
    uint8_t flags;
} y9HgzMZr_ctx_t;

void init_ctx(y9HgzMZr_ctx_t *ctx) {
    ctx->id = 3990;
    ctx->coefficient = 6.7;
    ctx->flags = 153;
}

int compute_214Wbq4(int input) {
    int result = input * 44;
    for(int i = 0; i < 10; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[87];
    double coefficient;
    uint8_t flags;
} Aq05x3X9_ctx_t;

void init_ctx(Aq05x3X9_ctx_t *ctx) {
    ctx->id = 5267;
    ctx->coefficient = 7.2;
    ctx->flags = 190;
}

typedef struct {
    uint32_t id;
    char name[32];
    double coefficient;
    uint8_t flags;
} Q3WmKrI5_ctx_t;

void init_ctx(Q3WmKrI5_ctx_t *ctx) {
    ctx->id = 1063;
    ctx->coefficient = 0.8;
    ctx->flags = 43;
}

int compute_IHECWm9(int input) {
    int result = input * 3;
    for(int i = 0; i < 5; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[91];
    double coefficient;
    uint8_t flags;
} 9TKhCmh7_ctx_t;

void init_ctx(9TKhCmh7_ctx_t *ctx) {
    ctx->id = 2776;
    ctx->coefficient = 1.2;
    ctx->flags = 154;
}

int compute_WcZiWtb(int input) {
    int result = input * 8;
    for(int i = 0; i < 6; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[64];
    double coefficient;
    uint8_t flags;
} HWYke2IZ_ctx_t;

void init_ctx(HWYke2IZ_ctx_t *ctx) {
    ctx->id = 7570;
    ctx->coefficient = 2.4;
    ctx->flags = 26;
}

int compute_Y0kDYPU(int input) {
    int result = input * 47;
    for(int i = 0; i < 7; i++) result += (i * 2);
    return result;
}

typedef struct {
    uint32_t id;
    char name[90];
    double coefficient;
    uint8_t flags;
} BcYjwQzu_ctx_t;

void init_ctx(BcYjwQzu_ctx_t *ctx) {
    ctx->id = 1072;
    ctx->coefficient = 0.2;
    ctx->flags = 247;
}

typedef struct {
    uint32_t id;
    char name[79];
    double coefficient;
    uint8_t flags;
} 8hW8g7So_ctx_t;

void init_ctx(8hW8g7So_ctx_t *ctx) {
    ctx->id = 7688;
    ctx->coefficient = 2.0;
    ctx->flags = 41;
}

typedef struct {
    uint32_t id;
    char name[73];
    double coefficient;
    uint8_t flags;
} JhzjUGZ2_ctx_t;

void init_ctx(JhzjUGZ2_ctx_t *ctx) {
    ctx->id = 9995;
    ctx->coefficient = 8.3;
    ctx->flags = 206;
}

int compute_mj4RbTj(int input) {
    int result = input * 57;
    for(int i = 0; i < 11; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[56];
    double coefficient;
    uint8_t flags;
} Z59tsBKl_ctx_t;

void init_ctx(Z59tsBKl_ctx_t *ctx) {
    ctx->id = 5596;
    ctx->coefficient = 0.3;
    ctx->flags = 105;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} KQTwdCfC_ctx_t;

void init_ctx(KQTwdCfC_ctx_t *ctx) {
    ctx->id = 4006;
    ctx->coefficient = 9.7;
    ctx->flags = 195;
}

int compute_zPMpGfP(int input) {
    int result = input * 66;
    for(int i = 0; i < 13; i++) result += (i * 1);
    return result;
}

typedef struct {
    uint32_t id;
    char name[75];
    double coefficient;
    uint8_t flags;
} UpfviPJv_ctx_t;

void init_ctx(UpfviPJv_ctx_t *ctx) {
    ctx->id = 1570;
    ctx->coefficient = 7.1;
    ctx->flags = 40;
}

typedef struct {
    uint32_t id;
    char name[83];
    double coefficient;
    uint8_t flags;
} 3RCh9XUv_ctx_t;

void init_ctx(3RCh9XUv_ctx_t *ctx) {
    ctx->id = 4067;
    ctx->coefficient = 5.2;
    ctx->flags = 123;
}

int compute_KjLIdkZ(int input) {
    int result = input * 29;
    for(int i = 0; i < 17; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[45];
    double coefficient;
    uint8_t flags;
} ww51NDG2_ctx_t;

void init_ctx(ww51NDG2_ctx_t *ctx) {
    ctx->id = 9861;
    ctx->coefficient = 1.9;
    ctx->flags = 169;
}

int compute_L1Oehbp(int input) {
    int result = input * 53;
    for(int i = 0; i < 18; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[55];
    double coefficient;
    uint8_t flags;
} NNYrpnYi_ctx_t;

void init_ctx(NNYrpnYi_ctx_t *ctx) {
    ctx->id = 9966;
    ctx->coefficient = 2.5;
    ctx->flags = 32;
}

typedef struct {
    uint32_t id;
    char name[86];
    double coefficient;
    uint8_t flags;
} X41teu1P_ctx_t;

void init_ctx(X41teu1P_ctx_t *ctx) {
    ctx->id = 5543;
    ctx->coefficient = 1.2;
    ctx->flags = 8;
}

int compute_QGrpZPs(int input) {
    int result = input * 16;
    for(int i = 0; i < 17; i++) result += (i * 5);
    return result;
}

typedef struct {
    uint32_t id;
    char name[68];
    double coefficient;
    uint8_t flags;
} jkyrZkdz_ctx_t;

void init_ctx(jkyrZkdz_ctx_t *ctx) {
    ctx->id = 3989;
    ctx->coefficient = 2.7;
    ctx->flags = 151;
}

typedef struct {
    uint32_t id;
    char name[39];
    double coefficient;
    uint8_t flags;
} oytr2zKv_ctx_t;

void init_ctx(oytr2zKv_ctx_t *ctx) {
    ctx->id = 7260;
    ctx->coefficient = 4.9;
    ctx->flags = 213;
}

typedef struct {
    uint32_t id;
    char name[71];
    double coefficient;
    uint8_t flags;
} 2KhGqNOV_ctx_t;

void init_ctx(2KhGqNOV_ctx_t *ctx) {
    ctx->id = 3391;
    ctx->coefficient = 5.0;
    ctx->flags = 56;
}

int compute_mUTs7TW(int input) {
    int result = input * 12;
    for(int i = 0; i < 13; i++) result += (i * 4);
    return result;
}

typedef struct {
    uint32_t id;
    char name[94];
    double coefficient;
    uint8_t flags;
} rLOyb7Zf_ctx_t;

void init_ctx(rLOyb7Zf_ctx_t *ctx) {
    ctx->id = 9866;
    ctx->coefficient = 0.9;
    ctx->flags = 80;
}

int compute_PAg3dfb(int input) {
    int result = input * 25;
    for(int i = 0; i < 3; i++) result += (i * 4);
    return result;
}

int compute_LqqMkoj(int input) {
    int result = input * 51;
    for(int i = 0; i < 10; i++) result += (i * 4);
    return result;
}

int compute_MQwb6HO(int input) {
    int result = input * 46;
    for(int i = 0; i < 11; i++) result += (i * 2);
    return result;
}

int compute_JfvcFv0(int input) {
    int result = input * 8;
    for(int i = 0; i < 16; i++) result += (i * 4);
    return result;
}

int compute_s3LzL4x(int input) {
    int result = input * 50;
    for(int i = 0; i < 14; i++) result += (i * 3);
    return result;
}

typedef struct {
    uint32_t id;
    char name[70];
    double coefficient;
    uint8_t flags;
} 0iLLgeYr_ctx_t;

void init_ctx(0iLLgeYr_ctx_t *ctx) {
    ctx->id = 7971;
    ctx->coefficient = 0.5;
    ctx->flags = 138;
}

typedef struct {
    uint32_t id;
    char name[34];
    double coefficient;
    uint8_t flags;
} it1Loxcd_ctx_t;

void init_ctx(it1Loxcd_ctx_t *ctx) {
    ctx->id = 4002;
    ctx->coefficient = 7.5;
    ctx->flags = 128;
}

typedef struct {
    uint32_t id;
    char name[65];
    double coefficient;
    uint8_t flags;
} GTjKSffg_ctx_t;

void init_ctx(GTjKSffg_ctx_t *ctx) {
    ctx->id = 4176;
    ctx->coefficient = 4.7;
    ctx->flags = 59;
}

typedef struct {
    uint32_t id;
    char name[39];
    double coefficient;
    uint8_t flags;
} unw82FKY_ctx_t;

void init_ctx(unw82FKY_ctx_t *ctx) {
    ctx->id = 1084;
    ctx->coefficient = 0.6;
    ctx->flags = 134;
}

